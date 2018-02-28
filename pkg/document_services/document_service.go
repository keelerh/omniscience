package document_services

import pb "github.com/keelerh/omniscience/protos"

type DocumentService interface {

	// Gets all documents from a given document hosting platform modified on or after the last modified time.
	GetAll(request *pb.GetAllDocumentsRequest, stream pb.GoogleDrive_GetAllServer) error
}