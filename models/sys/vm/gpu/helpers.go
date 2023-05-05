package gpu

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/goharbor/harbor/src/lib/log"
	"go-deploy/models"
	"go-deploy/models/dto/body"
	vm2 "go-deploy/models/sys/vm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


func (gpu *GPU) ToDto() body.GpuRead {
	id := base64.StdEncoding.EncodeToString([]byte(gpu.ID))

	var lease *body.GpuLease

	if gpu.Lease.VmID != "" {
		lease = &body.GpuLease{
			VmID: gpu.Lease.VmID,
			User: gpu.Lease.UserID,
			End:  gpu.Lease.End,
		}
	}

	return body.GpuRead{
		ID:    id,
		Name:  gpu.Data.Name,
		Lease: lease,
	}
}

func CreateGPU(id, host string, data GpuData) error {
	currentGPU, err := GetGpuByID(id)
	if err != nil {
		return err
	}

	if currentGPU != nil {
		return nil
	}

	gpu := GPU{
		ID:   id,
		Host: host,
		Lease: GpuLease{
			VmID:   "",
			UserID: "",
			End:    time.Time{},
		},
		Data: data,
	}

	_, err = models.GpuCollection.InsertOne(context.TODO(), gpu)
	if err != nil {
		err = fmt.Errorf("failed to create gpu. details: %s", err)
		return err
	}

	return nil
}

func GetGpuByID(id string) (*GPU, error) {
	var gpu GPU
	err := models.GpuCollection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&gpu)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		err = fmt.Errorf("failed to fetch gpu. details: %s", err)
		return nil, err
	}

	return &gpu, err
}

func GetAllGPUs(excludedHosts, excludedGPUs []string) ([]GPU, error) {
	if excludedHosts == nil {
		excludedHosts = make([]string, 0)
	}

	if excludedGPUs == nil {
		excludedGPUs = make([]string, 0)
	}

	filter := bson.D{
		{"host", bson.M{"$nin": excludedHosts}},
		{"id", bson.M{"$nin": excludedGPUs}},
	}

	var gpus []GPU
	cursor, err := models.GpuCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &gpus)
	if err != nil {
		return nil, err
	}

	return gpus, nil
}

func GetAllLeasedGPUs(excludedHosts, excludedGPUs []string) ([]GPU, error) {
	// return gpus that are leased
	if excludedHosts == nil {
		excludedHosts = make([]string, 0)
	}

	if excludedGPUs == nil {
		excludedGPUs = make([]string, 0)
	}

	// filter lease exist and vmId is not empty
	filter := bson.D{
		{"$and", []interface{}{
			bson.M{"lease.vmId": bson.M{"$ne": ""}},
			bson.M{"lease": bson.M{"$exists": true}},
		}},
		{"host", bson.M{"$nin": excludedHosts}},
		{"id", bson.M{"$nin": excludedGPUs}},
	}

	var gpus []GPU
	cursor, err := models.GpuCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &gpus)
	if err != nil {
		return nil, err
	}

	return gpus, nil
}

func GetAllAvailableGPUs(excludedHosts, excludedGPUs []string) ([]GPU, error) {
	if excludedHosts == nil {
		excludedHosts = make([]string, 0)
	}

	if excludedGPUs == nil {
		excludedGPUs = make([]string, 0)
	}

	filter := bson.D{
		{"$or", []interface{}{
			bson.M{"lease.vmId": ""},
			bson.M{"lease": bson.M{"$exists": false}},
		}},
		{"host", bson.M{"$nin": excludedHosts}},
		{"id", bson.M{"$nin": excludedGPUs}},
	}

	var gpus []GPU
	cursor, err := models.GpuCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &gpus)
	if err != nil {
		return nil, err
	}

	return gpus, nil
}

func AttachGPU(gpuID, vmID, user string, end time.Time) (bool, error) {
	vm, err := vm2.GetByID(vmID)
	if err != nil {
		return false, err
	}

	if vm == nil {
		return false, fmt.Errorf("vm not found")
	}

	gpu, err := GetGpuByID(gpuID)
	if err != nil {
		return false, err
	}

	if gpu == nil {
		return false, fmt.Errorf("gpu not found")
	}

	if gpu.Lease.VmID != "" && gpu.Lease.VmID != vmID {
		return false, fmt.Errorf("gpu is already attached to another vm")
	}

	if gpu.Lease.VmID == "" {
		filter := bson.D{
			{"id", gpuID},
			{"$or", []interface{}{
				bson.M{"lease.vmId": ""},
				bson.M{"lease": bson.M{"$exists": false}},
			}}}
		update := bson.M{
			"$set": bson.M{
				"lease.vmId": vmID,
				"lease.user": user,
				"lease.end":  end,
			},
		}

		opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

		err = models.GpuCollection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&gpu)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// this is not treated as an error, just another instance snatched the gpu before this one
				return false, nil
			}
			return false, err
		}
	}

	err = vm2.UpdateWithBsonByID(vmID, bson.D{{"gpuId", gpuID}})
	if err != nil {
		// remove lease, if this also fails, we are in a bad state...
		_, _ = models.GpuCollection.UpdateOne(
			context.TODO(),
			bson.D{{"id", gpuID}},
			bson.M{"$set": bson.M{"lease": GpuLease{}}},
		)
		log.Error("failed to remove lease after vm update failed. system is now in an inconsistent state. please fix manually. vm id:", vmID, " gpu id:", gpuID, ". details: %s", err)
		return false, err
	}

	return true, nil
}

func DetachGPU(vmID, userID string) (bool, error) {
	vm, err := vm2.GetByID(vmID)
	if err != nil {
		return false, err
	}

	if vm == nil {
		return false, fmt.Errorf("vm not found")
	}

	if vm.GpuID == "" {
		return true, nil
	}

	gpu, err := GetGpuByID(vm.GpuID)
	if err != nil {
		return false, err
	}

	if gpu == nil {
		return false, fmt.Errorf("gpu not found")
	}

	if gpu.Lease.VmID != vmID {
		return false, fmt.Errorf("vm is not attached to this gpu")
	}

	if gpu.Lease.UserID != userID {
		return false, fmt.Errorf("vm is not attached to this user")
	}

	filter := bson.D{
		{"id", gpu.ID},
		{"lease.vmId", vmID},
		{"lease.user", userID},
	}

	update := bson.M{
		"$set": bson.M{
			"lease.vmId": "",
			"lease.user": "",
			"lease.end":  time.Time{},
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = models.GpuCollection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&gpu)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// this is not treated as an error, just another instance snatched the gpu before this one
			return false, nil
		}
		return false, err
	}

	err = vm2.UpdateWithBsonByID(vmID, bson.D{{"gpuId", ""}})
	if err != nil {
		// remove lease, if this also fails, we are in a bad state...
		_, _ = models.GpuCollection.UpdateOne(
			context.TODO(),
			bson.D{{"id", gpu.ID}},
			bson.M{"$set": bson.M{"lease": GpuLease{}}},
		)
		log.Error("failed to remove lease after vm update failed. system is now in an inconsistent state. please fix manually. vm id:", vmID, " gpu id:", gpu.ID, ". details: %s", err)
	}

	return true, nil
}