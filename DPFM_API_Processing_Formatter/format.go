package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-planned-train-operation-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		RailwayLine:					data.RailwayLine,
		TrainOperationVersion:			data.TrainOperationVersion,
		WeekdayType:					data.WeekdayType,
		PlannedTrainOperationID:		data.PlannedTrainOperationID,
		ExpressType:					data.ExpressType,
		DepartureStation:				data.DepartureStation,
		DestinationStation:				data.DestinationStation,
		PlannedDepartureTime:			data.PlannedDepartureTime,
		PlannedArrivalTime:				data.PlannedArrivalTime,
		Description:					data.Description,
		OperationRemarks:				data.OperationRemarks,
		OperationCode:					data.OperationCode,
		IsSuspended:					data.IsSuspended,
		ValidityStartDate:				data.ValidityStartDate,
		ValidityEndDate:				data.ValidityEndDate,
		CreationDate:					data.CreationDate,
		CreationTime:					data.CreationTime,
		LastChangeDate:					data.LastChangeDate,
		LastChangeTime:					data.LastChangeTime,
		CreateUser:						data.CreateUser,
		LastChangeUser:					data.LastChangeUser,
		IsReleased:						data.IsReleased,
		IsMarkedForDeletion:			data.IsMarkedForDeletion,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item

	return &ItemUpdates{
		RailwayLine:             	dataHeader.RailwayLine,
		TrainOperationVersion:		data.TrainOperationVersion,
		WeekdayType:				data.WeekdayType,
		PlannedTrainOperationID:	data.PlannedTrainOperationID,
		RailwayLineStationID:		data.RailwayLineStationID,
		StopStation:				data.StopStation,
		PlannedPlatformNumber:		data.PlannedPlatformNumber,
		ExpressType:				data.ExpressType,
		PlannedArrivalTime:			data.PlannedArrivalTime,
		PlannedDepartureTime:		data.PlannedDepartureTime,
		Description:				data.Description,
		OperationRemarks:			data.OperationRemarks,
		IsSuspended:				data.IsSuspended,
		CreationDate:				data.CreationDate,
		CreationTime:				data.CreationTime,
		LastChangeDate:				data.LastChangeDate,
		LastChangeTime:				data.LastChangeTime,
		CreateUser:					data.CreateUser,
		LastChangeUser:				data.LastChangeUser,
		IsReleased:					data.IsReleased,
		IsMarkedForDeletion:		data.IsMarkedForDeletion,
	}
}
