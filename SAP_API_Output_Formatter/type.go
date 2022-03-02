package sap_api_output_formatter

type MeasurementDocument struct {
	 ConnectionKey           string `json:"connection_key"`
	 Result                  bool   `json:"result"`
	 RedisKey                string `json:"redis_key"`
	 Filepath                string `json:"filepath"`
	 APISchema               string `json:"api_schema"`
	 MeasurementDocument     string `json:"planned_order"`
	 Deleted                 bool   `json:"deleted"`
}

type Header struct {
	MeasurementDocument            string        `json:"MeasurementDocument"`
	MeasuringPoint                 string        `json:"MeasuringPoint"`
	MeasuringPointPositionNumber   string        `json:"MeasuringPointPositionNumber"`
	MsmtRdngDate                   string        `json:"MsmtRdngDate"`
	MsmtRdngTime                   string        `json:"MsmtRdngTime"`
	Characteristic                 string        `json:"Characteristic"`
	MsmtDocumentReferredOrder      string        `json:"MsmtDocumentReferredOrder"`
	RefdMaintOrderOpStatusObject   string        `json:"RefdMaintOrderOpStatusObject"`
	MaintenanceOrderOperation      string        `json:"MaintenanceOrderOperation"`
	MaintenanceOrderSubOperation   string        `json:"MaintenanceOrderSubOperation"`
	MsmtIsDoneAfterTaskCompltn     bool          `json:"MsmtIsDoneAfterTaskCompltn"`
	CharcValueUnit                 string        `json:"CharcValueUnit"`
	MeasurementReading             float64       `json:"MeasurementReading"`
	MeasurementReadingInEntryUoM   float64       `json:"MeasurementReadingInEntryUoM"`
	MeasurementReadingEntryUoM     string        `json:"MeasurementReadingEntryUoM"`
	MeasurementCounterReading      float64       `json:"MeasurementCounterReading"`
	MsmtCounterReadingDifference   float64       `json:"MsmtCounterReadingDifference"`
	TotalMsmtRdngIsSetExternally   bool          `json:"TotalMsmtRdngIsSetExternally"`
	MeasuringPointTargetValue      int           `json:"MeasuringPointTargetValue"`
	MsmtValuationCode              string        `json:"MsmtValuationCode"`
	MeasurementDocumentText        string        `json:"MeasurementDocumentText"`
	MeasurementDocumentHasLongText bool          `json:"MeasurementDocumentHasLongText"`
	MsmtRdngByUser                 string        `json:"MsmtRdngByUser"`
	MsmtRdngStatus                 string        `json:"MsmtRdngStatus"`
	MsmtCntrReadingDiffIsEntered   bool          `json:"MsmtCntrReadingDiffIsEntered"`
	MsmtRdngIsReversed             bool          `json:"MsmtRdngIsReversed"`
	MsmtCounterReadingIsReplaced   bool          `json:"MsmtCounterReadingIsReplaced"`
}
