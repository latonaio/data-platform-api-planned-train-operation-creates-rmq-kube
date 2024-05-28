package dpfm_api_processing_formatter

type HeaderUpdates struct {
	RailwayLine					int		`json:"RailwayLine"`
	TrainOperationVersion		int		`json:"TrainOperationVersion"`
	WeekdayType					string	`json:"WeekdayType"`
	PlannedTrainOperationID		int		`json:"PlannedTrainOperationID"`
	ExpressType					string	`json:"ExpressType"`
	DepartureStation			int		`json:"DepartureStation"`
	DestinationStation			int		`json:"DestinationStation"`
	PlannedDepartureTime		string	`json:"PlannedDepartureTime"`
	PlannedArrivalTime			string	`json:"PlannedArrivalTime"`
	Description					string	`json:"Description"`
	OperationRemarks			*string	`json:"OperationRemarks"`
	OperationCode				*string	`json:"OperationCode"`
	IsSuspended					*bool	`json:"IsSuspended"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	LastChangeUser				int		`json:"LastChangeUser"`
}

type ItemUpdates struct {
	RailwayLine					int		`json:"RailwayLine"`
	TrainOperationVersion		int		`json:"TrainOperationVersion"`
	WeekdayType					string	`json:"WeekdayType"`
	PlannedTrainOperationID		int		`json:"PlannedTrainOperationID"`
	RailwayLineStationID		int		`json:"RailwayLineStationID"`
	StopStation					int		`json:"StopStation"`
	PlannedPlatformNumber		string	`json:"PlannedPlatformNumber"`
	ExpressType					string	`json:"ExpressType"`
	PlannedArrivalTime			string	`json:"PlannedArrivalTime"`
	PlannedDepartureTime		*string	`json:"PlannedDepartureTime"`
	Description					string	`json:"Description"`
	OperationRemarks			*string	`json:"OperationRemarks"`
	IsSuspended					*bool	`json:"IsSuspended"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	LastChangeUser				int		`json:"LastChangeUser"`
}

