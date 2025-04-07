package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tender struct {
	ID                  primitive.ObjectID      `bson:"_id,omitempty"`
	Title               string                  `bson:"title"`
	Subtitle            string                  `bson:"subtitle"`
	NoticeNumber        string                  `bson:"notice_number"`
	ObjectName          string                  `bson:"object_name"`
	Method              string                  `bson:"procurement_method"`
	EPlatform           ElectronicPlatform      `bson:"electronic_platform"`
	PlacingOrganization string                  `bson:"placing_organization"`
	Contact             ContactInfo             `bson:"contact"`
	Procedure           ProcedureInfo           `bson:"procedure"`
	Contract            ContractInfo            `bson:"contract"`
	Execution           ExecutionInfo           `bson:"execution"`
	DeliveryPlace       string                  `bson:"delivery_place"`
	Items               []ProcurementItem       `bson:"items"`
	Security            ApplicationSecurityInfo `bson:"security"`
}
