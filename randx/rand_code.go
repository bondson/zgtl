/**
  @Description: 随机码生成
  @Author: ZPS
**/

package randx

import (
	"math/rand"
)

type TYPE int

const (
	TYPE_DIGIT   TYPE = 1 //数字//
	TYPE_LETTER  TYPE = 2 //小写字母
	TYPE_CAPITAL TYPE = 3 //大写字母
	TYPE_MIXED   TYPE = 4 //数字+字母混合
)

// RandCode 根据传入的长度和类型生成随机字符串,这个方法目前可以生成数字、字母、数字+字母的随机字符串
// @param length: 生成字符串的长度
// @param typ: 生成字符串的类型: 默认数字 TYPE_DIGIT
func RandCode(length int, typ ...TYPE) string {
	if len(typ) > 0 {
		switch typ[0] {
		case TYPE_LETTER:
			return generate("abcdefghijklmnopqrstuvwxyz", length, 5)
		case TYPE_CAPITAL:
			return generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ", length, 5)
		case TYPE_MIXED:
			return generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length, 7)
		}
	}
	//TYPE_DIGIT
	return generate("0123456789", length, 4)
}

// generate 根据传入的随机源和长度生成随机字符串,一次随机，多次使用
func generate(source string, length, idxBits int) string {

	//掩码
	//例如： 使用低6位：0000 0000 --> 0011 1111
	idxMask := 1<<idxBits - 1

	// 63位最多可以使用多少次
	remain := 63 / idxBits

	//cache 随机位缓存
	cache := rand.Int63()

	result := make([]byte, length)

	for i := 0; i < length; {
		//如果使用次数剩余0，重新获取随机
		if remain == 0 {
			cache, remain = rand.Int63(), 63/idxBits
		}

		//利用掩码获取有效的随机数位
		if randIndex := int(cache & int64(idxMask)); randIndex < len(source) {
			result[i] = source[randIndex]
			i++
		}

		//使用下一组随机位
		cache >>= idxBits

		//扣减remain
		remain--

	}
	return string(result)

}
