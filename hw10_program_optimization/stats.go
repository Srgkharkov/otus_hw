package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type User struct {
	//ID       int
	//Name     string
	//Username string
	Email string `json:"Email"`
	//Phone    string
	//Password string
	//Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {

	fscanner := bufio.NewScanner(r)
	i := 0
	for fscanner.Scan() {
		var user User
		if err = user.UnmarshalJSON(fscanner.Bytes()); err != nil {
			return
		}
		result[i] = user
		i++

	}

	//content, err := ioutil.ReadAll(r)
	//if err != nil {
	//	return
	//}
	//
	//lines := bytes.Split(content, []byte("\n"))
	////lines := strings.Split(string(content), "\n")
	//for i, line := range lines {
	//	var user User
	//	if err = user.UnmarshalJSON([]byte(line)); err != nil {
	//		//json.Unmarshal([]byte(line), &user); err != nil {
	//		return
	//	}
	//	//if err = json.Unmarshal([]byte(line), &user); err != nil {
	//	//	return
	//	//}
	//	result[i] = user
	//}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	//r, err :=
	domain = "." + domain
	lendomain := len(domain)

	for _, user := range u {
		li := strings.LastIndex(user.Email, domain)
		if li == -1 {
			continue
		}
		if li != len(user.Email)-lendomain {
			continue
		}
		subdomain := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
		num := result[subdomain]
		num++
		result[subdomain] = num
		//matched := r.Match([]byte(user.Email))
		//matched := r.MatchString(user.Email)

		//matched, err := regexp.Match("\\."+domain, []byte(user.Email))
		//if err != nil {
		//	return nil, err
		//}
		//
		//if matched {
		//	num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
		//	num++
		//	result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		//}
	}
	return result, nil
}
