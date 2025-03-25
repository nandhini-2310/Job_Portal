package services

import (
	db "DB_GORM/DB"
	"DB_GORM/models"
	pb "DB_GORM/pb_file"
	"DB_GORM/utils"
	"context"
	"fmt"
)

type Job struct {
	pb.UnimplementedJobServiceServer
}

func (j *Job) CreateJob(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
	var recruiter models.Recruiter
	if err := db.DB.First(&recruiter, req.RecruiterId).Error; err != nil {
		utils.ErrorLog.Println("Recruiter Not Found", err)
		return nil, fmt.Errorf("recruiter with ID %d not found", req.RecruiterId)
	}
	job := models.Job{
		RecruiterID:    uint(req.RecruiterId),
		Title:          req.Title,
		Description:    req.Description,
		SkillsRequired: req.SkillsRequired,
		Location:       req.Location,
		Salary:         req.Salary,
		JobType:        req.JobType,
	}
	data := db.DB.Create(&job)

	if data.Error != nil {
		utils.ErrorLog.Println("Error in Job Creation", data.Error)
		return nil, data.Error
	}
	return &pb.JobResponse{
		Message: "Job Created Successfully!!",
	}, nil
}

func (j *Job) GetJob(ctx context.Context, req *pb.JobID) (*pb.JobRequest, error) {
	var job models.Job
	data := db.DB.First(&job, req.Id)
	if data.Error != nil {
		utils.ErrorLog.Println("Job Not Found", data.Error)
		return nil, data.Error
	}

	return &pb.JobRequest{
		Title:          job.Title,
		Description:    job.Description,
		SkillsRequired: job.SkillsRequired,
		Location:       job.Location,
		Salary:         job.Salary,
		JobType:        job.JobType,
	}, nil
}

func (j *Job) DeleteJob(ctx context.Context, req *pb.JobID) (*pb.JobResponse, error) {
	var job models.Job

	data := db.DB.First(&job, req.Id)
	if data.Error != nil {
		utils.ErrorLog.Println("Job Not Found", data.Error)
		return nil, data.Error
	}

	data1 := db.DB.Delete(&job)
	if data1.Error != nil {
		utils.ErrorLog.Println("Job Not able to Delete", data1.Error)
		return nil, data.Error
	}

	return &pb.JobResponse{
		Message: "Job Deleted Successfully...",
	}, nil
}
