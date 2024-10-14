package nets

import "github.com/gin-gonic/gin"
type SudoAuthentication struct{
	SudoName      string `json:"sudoName"`
	SudoPassword  string `json:"sudoPassword"`
}
func authenticateSudo(s SudoAuthentication)(bool){
	if 
}

func sudoRegisterAdmin(c *gin.Context) {
	type registerAdminInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		UserName      string `json:"userName"`
		AdminAccount  string `json:"adminAccount"`
		AdminPassword string `json:"adminPassword"`
	}
	var getData registerAdminInfo
	if err := c.ShouldBindJSON(&getData); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

}
