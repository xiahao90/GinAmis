package utils
//一些公共函数
import (
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "time"
    "strings"
    "github.com/dgrijalva/jwt-go"
    // "errors"
    // "fmt"
)
func Sha256(input string)(string){
    // 计算 SHA-256 哈希值
    hashInBytes := sha256.Sum256([]byte(input))
    // 将哈希值转换为十六进制字符串
    hashString := hex.EncodeToString(hashInBytes[:])
    return hashString
}

func Base64decode(base64String string)(string,error){
    // 使用标准库的解码函数进行解码
    decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
    if err != nil {
        return "",err
    }
    // 将解码后的字节转换为字符串并输出
    password := string(decodedBytes)
    return password,nil
}
//计算获取jwt
//secret-key密钥
//extime过期时间（小时）
func JetEncode(uid int32,name string,superadmin int,role []string,rolename string,secretkey string,extime int) (string, error) {
    var secretKey = []byte(secretkey) // 自定义密钥，请根据实际情况修改
    // 创建一个新的 Token
    token := jwt.New(jwt.SigningMethodHS256)
    // 设置声明（Claims）
    claims := token.Claims.(jwt.MapClaims)
    claims["uid"] = uid
    claims["name"] = name
    claims["superadmin"] = superadmin
    claims["role"] = strings.Join(role, ";")+";"
    claims["rolename"] = rolename
    claims["exp"] = time.Now().Add(time.Hour * time.Duration(extime)).Unix() // Token 过期时间为 1 小时
    // 签名 Token
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}
//从jwt获取数据
func JetDecode(tokenString string,secretkey string,) (*jwt.Token, error) {
    // 解析 Token
    var secretKey = []byte(secretkey) // 自定义密钥，请根据实际情况修改
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return nil, err
    }
    return token, nil
}
// getMapDefule 获取 map 中的值或默认值
func GetMapDefule(myMap map[string]interface{}, key string, defaultValue interface{}) interface{} {
    if value, ok := myMap[key]; ok {
        return value
    }
    return defaultValue
}

// // 验证map中参数是否存在
// func CheckJson(data map[string]interface{}, params []string) (error) {
//     // var missingParams []string
//     for _, param := range params {
//         if _, exists := data[param]; !exists {
//             return errors.New(param+"必传")
//             // missingParams = append(missingParams, param)
//         }
//     }
//     return nil
//     // return missingParams
// }
