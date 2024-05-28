package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-planned-train-operation-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-planned-train-operation-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-planned-train-operation-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *subfuncSDC.Message.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		RailwayLine:                    input.Header.RailwayLine,
		TrainOperationVersion:			input.Header.TrainOperationVersion,
		WeekdayType:					input.Header.WeekdayType,
		PlannedTrainOperationID:		*input.Header.PlannedTrainOperationID,
		ExpressType:					input.Header.ExpressType,
		DepartureStation:				input.Header.DepartureStation,
		DestinationStation:				input.Header.DestinationStation,
		PlannedDepartureTime:			input.Header.PlannedDepartureTime,
		PlannedArrivalTime:				input.Header.PlannedArrivalTime,
		Description:					input.Header.Description,
		OperationRemarks:				input.Header.OperationRemarks,
		OperationCode:					input.Header.OperationCode,
		IsSuspended:					input.Header.IsSuspended,
		ValidityStartDate:				input.Header.ValidityStartDate,
		ValidityEndDate:				input.Header.ValidityEndDate,
		CreationDate:					input.Header.CreationDate,
		CreationTime:					input.Header.CreationTime,
		LastChangeDate:					input.Header.LastChangeDate,
		LastChangeTime:					input.Header.LastChangeTime,
		CreateUser:						input.Header.CreateUser,
		LastChangeUser:					input.Header.LastChangeUser,
		IsReleased:						input.Header.IsReleased,
		IsMarkedForDeletion:			input.Header.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToItem(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var items []sub_func_complementer.Item

	items = append(
		items,
		sub_func_complementer.Item{
			RailwayLine:          		input.Header.RailwayLine,
			TrainOperationVersion:		input.Header.Item[0].TrainOperationVersion,
			WeekdayType:				input.Header.Item[0].WeekdayType,
			PlannedTrainOperationID:	input.Header.Item[0].PlannedTrainOperationID,
			RailwayLineStationID:		input.Header.Item[0].RailwayLineStationID,
			StopStation:				input.Header.Item[0].StopStation,
			PlannedPlatformNumber:		input.Header.Item[0].PlannedPlatformNumber,
			ExpressType:				input.Header.Item[0].ExpressType,
			PlannedArrivalTime:			input.Header.Item[0].PlannedArrivalTime,
			PlannedDepartureTime:		input.Header.Item[0].PlannedDepartureTime,
			Description:				input.Header.Item[0].Description,
			OperationRemarks:			input.Header.Item[0].OperationRemarks,
			IsSuspended:				input.Header.Item[0].IsSuspended,
			CreationDate:				input.Header.Item[0].CreationDate,
			CreationTime:				input.Header.Item[0].CreationTime,
			LastChangeDate:				input.Header.Item[0].LastChangeDate,
			LastChangeTime:				input.Header.Item[0].LastChangeTime,
			CreateUser:					input.Header.Item[0].CreateUser,
			LastChangeUser:				input.Header.Item[0].LastChangeUser,
			IsReleased:					input.Header.Item[0].IsReleased,
			IsMarkedForDeletion:		input.Header.Item[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.Item = &items

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
