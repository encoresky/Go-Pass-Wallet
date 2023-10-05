package internal

import "net/http"

func GetListUpdatablePasses(w http.ResponseWriter, r *http.Request) {
	// Extract the deviceLibraryIdentifier from the parameter request.
	// Extract the passTypeIdentifier from the parameter request
	// Extract the previousLastUpdated from the query, The value of the lastUpdated key from the SerialNumbers object returned in a previous request. This value limits the results of the current request to the passes updated since that previous request.

	// implement logic to get list of updatable passes and send the passes serialNumber in array
}
