package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_METADATA = "/config/metadata"
const URL_ROUTER_METADATA_BY_ID = "/config/metadata/:metadata-id"
const URL_ROUTER_METADATA_QUERY = "/config/metadata/query"
const URL_ROUTER_METADATA_ALL = "/config/metadata/all"

func MetadataRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_METADATA, config.CreateMetadata)
	r.GET(URL_ROUTER_METADATA_BY_ID, config.RetrieveMetadata)
	r.PUT(URL_ROUTER_METADATA, config.UpdateMetadata)
	r.DELETE(URL_ROUTER_METADATA_BY_ID, config.DeleteMetadata)
	r.PUT(URL_ROUTER_METADATA_BY_ID, config.DeleteMetadataWithFlag)
	r.POST(URL_ROUTER_METADATA_QUERY, config.QueryMetadata)
	r.GET(URL_ROUTER_METADATA_ALL, config.GetAllMetadatas)
	return r
}