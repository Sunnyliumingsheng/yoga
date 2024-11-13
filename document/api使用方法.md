# api
there are the usage of these api for weixin app

1. login
there are two api /api/login and /api/register   
the first one receives the json
    type authenticationInfo struct {
            Session string `json:"session"`
            Token   string `json:"token"`
        }
and if the first one feedback is error,it indicates that this is a new user or the user has not logged for a long time , you should use the second api ,and it recieves the json   
    type codeStruct struct {
            Code string `json:"code"`
        }
besides ,please use localStorage save the token and session   
2. name
for a new user ,it will get a default name ,besides ,name is very important and unique.   
attention please ,name and nickname is not one attribute, you should remind user that please use the realname as 'name' . nickname usually is the wechat name    
the rename api recieve a nested json:
    type userInfo struct {
            NewName        string             `json:"newName"`
            Authentication AuthenticationInfo `json:"authentication"`
        }
    type AuthenticationInfo struct {
        Session string `json:"session"`
        Token   string `json:"token"`
    }
