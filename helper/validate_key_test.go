package helper

import (
	"testing"
)

func TestValidateRsaPrivateKeySucceeded(t *testing.T) {
	goodPrivateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA5jBjuJNj2dTTr7l3mkz5+uxe8bMmslGzA571zAAfQ3g2QBzY\nJahYCNzjSlgGnHpmBtBAwlXF9baxsgnANUKf9FguQrsBakX4aiC8RQACTVgjVRgv\nllNfgv82qvv4FrrGKpsV+5hELmg+gSLKR1I5dCGrrLb1hsBofupN/Vtw5hQjbH3P\nErFaF7VuC1ozokwV3utrDtpm+EoTO/dufvQwxc7RkBhkfGHVrARaO6GMWaT2uFsp\n0VJVb1xADEKSH5QgZRsSJquJz7o76QUK1Y/e1k1vmvNHA7+3rbeXHO+Mc6f7HbOv\ns+v3P/nlqDsoqE47EdKDjf9wVpczK+nmuqWl6wIDAQABAoIBABDNJGIzA91A7wiK\n+YNeLfrWKOHsaR0n4HCZuCghNCb5DcClOlOZU3+mG8Lv5N+kAXFq3ucwWogiQxsT\nIh/hEorDqvC89I/8tnKMnCFPLFvX3JBX+RU4VabamPNm+8cw78jpU/Puu79FZZaI\nYwpMopvq3zx2iDpDLf1hfbrJ41Yvq6senN5Y0SPZzwChZarUHtC3UzhE/rmRTDo8\ngBTQdnKY75jYJbGYnTmtgkjd6fk+Gjp0qb8SDR9OL7vDYB8t6uMHlLAqCYsJr8Ra\nXQti6zCvYYJGXpQjFFZgbXp5n5imW1mB2aqNg9IAIA8p8olbzPt4fO5p8+scizbG\nSYV7MbUCgYEA7ssQJa6tfyDJzm/aD3N6zqYdefkR7SYznmiU+JKgYcoYpqc6sbjP\n+u2WeA/TuTSnfHDWu6eZZwjc+ZZK7fPf6IrZSjWYHR/Hk+7RDw2BjeYQsHnn+MDh\nF/i//6N1Vy+0BWO9zeFkR/SOT6XFj+6/DuWGKdEctwRNOSW9h2kGpPUCgYEA9sab\nm18OBjlrHAnU80hiJE+D2bhACKesjnxF5GSNN/RHZzVo0IGsaJpECPN3ZyNhbDpb\nnNFLvuJ68DGRAFMXMeNiJbe4U7erOJQ4B5GTyMITJz0X/0TyxkS6KV4tmB/AQx6k\nf/HGQNWGSbSJrn1w+HX3ZaaTO6HzlRLNqFsHU18CgYEAqxgp6KYAv77Rea1g04NN\ncbVKF3PjaUTe5VhPrM76RTbVdMsLXf1qX4NONZY5gFD+1EXRRcFvQ4jxM4+A/n+n\nZeneiLJf3DZB/He3qQ5nTjGSsa+XCv/ACDn973/B+oc+eEIf49zjyj6qnNzM8jUB\nBf5ko5+l8GgPoKuu4c/Zp+UCgYEA0z3HzgRcWO+lxGaWJ3r6p+J5F1IlXkNJ7t5q\nZcu9JvywcqTBiFq4XwJO4uqmd19N8fsymNaAZHIykpGbcg+ud0IRrf/Khb9fjhsy\nqyfuvZeEFb6yYA2BVD9YrDi2BtaHGe2NGwi3kKA3R7iHcxpmLgmtMRKaTh2gCSgD\na+4gIzkCgYEAjBB+tHIDYxyA5RKzpst8rPqTLbK1dcGXrmE2uoPDhAh70XjsGwUq\nyAc/lRzNEK//2EmKJGuU1EW9CbGZ/kgJFq9wogMXULb1QMQSKkTVpm9rnXiTAzER\nySG4w3L9efhC2GHYs6XvW7lZxyJELUHN4AqyULfri/bsVz+Ekumzvq4=\n-----END RSA PRIVATE KEY-----"
	err := ValidateRsaPrivateKey(goodPrivateKey)
	if err != nil {
		t.Fatalf("The function doesn't succeed with good private key provided: %v", err)
	}
}

func TestValidateRsaPrivateKeyFailed(t *testing.T) {
	badPrivateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpIBAAKCAQEA5jBjuJNj2dTTr7l3mkz5+uxe8bMmslGzA571zAAfQ3g2QBzY\nJahYCNzjSlgGnHpmBtBAwlXF9baxsgnANUKf9FguQrsBakX4aiC8RQACTVgjVRgv\nllNfgv82qvv4FrrGKpsV+5hELmg+gSLKR1I5dCGrrLb1hsBofupN/Vtw5hQjbH3P\nErFaF7VuC1ozokwV3utrDtpm+EoTO/dufvQwxc7RkBhkfGHVrARaO6GMWaT2uFsp\n0VJVb1xADEKSH5QgZRsSJquJz7o76QUK1Y/e1k1vmvNHA7+3rbeXHO+Mc6f7HbOv\ns+v3P/nlqDsoqE47EdKDjf9wVpczK+nmuqWl6wIDAQABAoIBABDNJGIzA91A7wiK\n+YNeLfrWKOHsaR0n4HCZuCghNCb5DcClOlOZU3+mG8Lv5N+kAXFq3ucwWogiQxsT\nIh/hEorDqvC89I/8tnKMnCFPLFvX3JBX+RU4VabamPNm+8cw78jpU/Puu79FZZaI\nYwpMopvq3zx2iDpDLf1hfbrJ41Yvq6senN5Y0SPZzwChZarUHtC3UzhE/rmRTDo8\ngBTQdnKY75jYJbGYnTmtgkjd6fk+Gjp0qb8SDR9OL7vDYB8t6uMHlLAqCYsJr8Ra\nXQti6zCvYYJGXpQjFFZgbXp5n5imW1mB2aqNg9IAIA8p8olbzPt4fO5p8+scizbG\nSYV7MbUCgYEA7ssQJa6tfyDJzm/aD3N6zqYdefkR7SYznmiU+JKgYcoYpqc6sbjP\n+u2WeA/TuTSnfHDWu6eZZwjc+ZZK7fPf6IrZSjWYHR/Hk+7RDw2BjeYQsHnn+MDh\nF/i//6N1Vy+0BWO9zeFkR/SOT6XFj+6/DuWGKdEctwRNOSW9h2kGpPUCgYEA9sab\nm18OBjlrHAnU80hiJE+D2bhACKesjnxF5GSNN/RHZzVo0IGsaJpECPN3ZyNhbDpb\nnNFLvuJ68DGRAFMXMeNiJbe4U7erOJQ4B5GTyMITJz0X/0TyxkS6KV4tmB/AQx6k\nf/HGQNWGSbSJrn1w+HX3ZaaTO6HzlRLNqFsHU18CgYEAqxgp6KYAv77Rea1g04NN\ncbVKF3PjaUTe5VhPrM76RTbVdMsLXf1qX4NONZY5gFD+1EXRRcFvQ4jxM4+A/n+n\nZeneiLJf3DZB/He3qQ5nTjGSsa+XCv/ACDn973/B+oc+eEIf49zjyj6qnNzM8jUB\nBf5ko5+l8GgPoKuu4c/Zp+UCgYEA0z3HzgRcWO+lxGaWJ3r6p+J5F1IlXkNJ7t5q\nZcu9JvywcqTBiFq4XwJO4uqmd19N8fsymNaAZHIykpGbcg+ud0IRrf/Khb9fjhsy\nqyfuvZeEFb6yYA2BVD9YrDi2BtaHGe2NGwi3kKA3R7iHcxpmLgmtMRKaTh2gCSgD\na+4gIzkCgYEAjBB+tHIDYxyA5RKzpst8rPqTLbK1dcGXrmE2uoPDhAh70XjsGwUq\nyAc/lRzNEK//2EmKJGuU1EW9CbGZ/kgJFq9wogMXULb1QMQSKkTVpm9rnXiTAzER\nySG4w3L9efhC2GHYs6XvW7lZxyJELUHN4AqyULfri/bsVz+Ekumzvq4=\n-----END RSA PRIVATE KEY-----"
	err := ValidateRsaPrivateKey(badPrivateKey)
	if err == nil {
		t.Fatalf("Although the bad private key provided, the function does not fail")
	}
}

func TestValidateRsaPublicKeySucceeded(t *testing.T) {
	goodPublicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsdZVUMlKo9X/A3Uyttch\nFrxYnBiyi3R/DhzuUW5A1gamQwkQW5DG/hPfJE4+yuH9f4k7H+0mzCqHNgjhXndZ\n6sL+l3SO4+nu9QMi91oNZ4NIQagUuC1js4Va8t8/LKuW20nUklX8B6tnYAc0m4bj\nGDlYvLmG7vjUbmv2jnnDkYSxBYrIUCWODXhtzP9Uh2o3V8ggZusUbnFr1YBLRKaT\nvuTTiJdRVJC+gLYloyJ1EA4hXK/o0r1VmFv6z8GlOxd8T1zKLjNQY76u8eELIIU9\nPhOFV2uv7ipweyfWVCjBaiEdMcWBWRz9IhCutrXp7zuUW2WM3yi+HY6T4DgOGsdZ\nQQIDAQAB\n-----END PUBLIC KEY-----"
	err := ValidateRsaPublicKey(goodPublicKey)
	if err != nil {
		t.Fatalf("The function doesn't succeed with good public key provided: %v", err)
	}
}

func TestValidateRsaPublicKeyFailed(t *testing.T) {
	badPublicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCQ8AMIIBCgKCAQEAsdZVUMlKo9X/A3Uyttch\nFrxYnBiyi3R/DhzuUW5A1gamQwkQW5DG/hPfJE4+yuH9f4k7H+0mzCqHNgjhXndZ\n6sL+l3SO4+nu9QMi91oNZ4NIQagUuC1js4Va8t8/LKuW20nUklX8B6tnYAc0m4bj\nGDlYvLmG7vjUbmv2jnnDkYSxBYrIUCWODXhtzP9Uh2o3V8ggZusUbnFr1YBLRKaT\nvuTTiJdRVJC+gLYloyJ1EA4hXK/o0r1VmFv6z8GlOxd8T1zKLjNQY76u8eELIIU9\nPhOFV2uv7ipweyfWVCjBaiEdMcWBWRz9IhCutrXp7zuUW2WM3yi+HY6T4DgOGsdZ\nQQIDAQAB\n-----END PUBLIC KEY-----"
	err := ValidateRsaPublicKey(badPublicKey)
	if err == nil {
		t.Fatalf("Failed: although the bad public key provided, the function does not fail")
	}
}
