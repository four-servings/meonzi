package kakao

import (
	"fmt"
	"time"
)

type ErrorCode int64

const (
	// HTTP STATUS 400~403
	WrongParam ErrorCode = -2
	// HTTP STATUS 400~403
	WrongTokenOrInvalidate ErrorCode = -401

	// HTTP STATUS 503
	ServiceInspection ErrorCode = -9798
)

type Error struct {
	Code ErrorCode `json:"code"`
	Msg  string    `json:"msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("code:%d/message:%s", e.Code, e.Msg)
}

type AgeRangeType string

const (

	// 1세 이상 10세 미만
	Age1To9Range AgeRangeType = "1~9"

	// 10세 이상 15세 미만
	Age10To14Range AgeRangeType = "10~14"

	// 15세 이상 20세 미만
	Age15To19Range AgeRangeType = "15~19"

	// 20세 이상 30세 미만
	Age20To29Range AgeRangeType = "20~29"

	// 30세 이상 40세 미만
	Age30To39Range AgeRangeType = "30~39"

	// 40세 이상 50세 미만
	Age40To49Range AgeRangeType = "40~49"

	// 50세 이상 60세 미만
	Age50To59Range AgeRangeType = "50~59"

	// 60세 이상 70세 미만
	Age60To69Range AgeRangeType = "60~69"

	// 70세 이상 80세 미만
	Age70To79Range AgeRangeType = "70~79"

	// 80세 이상 90세 미만
	Age80To89Range AgeRangeType = "80~89"

	// 90세 이상
	Age90EqThanRange AgeRangeType = "90~"
)

type BirthdayType string

const (
	SolarType BirthdayType = "SOLAR"
	LunarType BirthdayType = "LUNAR"
)

type GenderType string

const (
	FemaleType GenderType = "female"
	MaleType   GenderType = "male"
)

type UserProfile struct {

	// properties = 사용자가 해당 서비스에서 설정한 닉네임, 기본 값은 카카오계정 닉네임
	// profile = 닉네임
	Nickname string `json:"nickname"`

	// properties = 사용자가 해당 서비스에서 설정한 프로필 이미지 URL, 기본 값은 카카오계정 프로필 이미지 URL
	// profile = 프로필 이미지 URL, 640px * 640px 또는 480px * 480px
	ProfileImage string `json:"profile_image"`

	// properties = 사용자가 해당 서비스에서 설정한 프로필 이미지 썸네일 URL, 기본 값은 카카오계정 썸네일 URL
	// profile = 프로필 미리보기 이미지 URL, 110px * 110px 또는 100px * 100px
	ThumbnailImageUrl string `json:"thumbnail_image_url"`
}

type UserInfo struct {

	// 회원번호
	Id int64 `json:"id"`

	// 카카오계정 정보
	KakaoAccount struct {
		Profile UserProfile `json:"profile"`

		// 사용자 동의 시 프로필 제공 가능
		// https://developers.kakao.com/docs/latest/ko/kakaologin/common#needs_agreement
		ProfileNeedsAgreement bool `json:"profile_needs_agreement"`

		// TODO description
		HasEmail bool `json:"has_email"`

		// 사용자 동의 시 이메일 제공 가능
		EmailNeedsAgreement bool `json:"email_needs_agreement"`

		// 이메일이 다른 카카오계정에 사용돼 만료되었다면 무효(false)
		IsEmailValid bool `json:"is_email_valid"`

		// 이메일 인증 여부
		IsEmailVerified bool `json:"is_email_verified"`

		// 대표 이메일
		Email string `json:"email"`

		// 사용자 동의 시 연령대 제공 가능
		AgeRangeNeedsAgreement bool `json:"age_range_needs_agreement"`

		// 연령대
		AgeRange AgeRangeType `json:"age_range"`

		// TODO description
		HasBirthday bool `json:"has_birthday"`

		// 사용자 동의 시 생일 제공 가능
		BirthdayNeedsAgreement bool `json:"birthday_needs_agreement"`

		// 생일, MMDD 형식
		Birthday string `json:"birthday"`

		// 생일 양력/음력 구분, 양력(SOLAR)/음력(LUNAR)
		BirthdayType BirthdayType `json:"birthday_type"`

		// TODO description
		HasGender bool `json:"has_gender"`

		// 사용자 동의 시 성별 제공 가능
		GenderNeedsAgreement bool `json:"gender_needs_agreement"`

		// 성별, female/male
		Gender string `json:"gender"`

		// TODO description, 존재하는지 확실치 않음
		HasPhoneNumber bool `json:"has_phone_number"`

		// 사용자 동의 시 전화번호 제공 가능
		PhoneNumberNeedsAgreement bool `json:"phone_number_needs_agreement"`

		// 전화번호, +00 00-0000-0000 또는 +00 00 0000 0000 형식, 국가마다 하이픈(-) 위치나 값 다를 수 있음
		// https://github.com/google/libphonenumber
		PhoneNumber string `json:""`
	} `json:"kakao_account"`

	// 서비스에 연결 완료된 시각, UTC (RFC3339 internet date/time format)
	ConnectedAt time.Time `json:"connected_at"`

	// 사용자 프로퍼티(Property)
	// https://developers.kakao.com/docs/latest/ko/kakaologin/prerequisite#user-properties
	Properties UserProfile `json:"properties"`
}

/*
{
    "id": 1258007123,
    "connected_at": "2020-01-28T10:59:18Z",
    "properties": {
        "nickname": "이재성",
        "profile_image": "-",
        "thumbnail_image": "-"
    },
    "kakao_account": {
        "profile_needs_agreement": false,
        "profile": {
            "nickname": "이재성",
            "thumbnail_image_url": "-",
            "profile_image_url": "-"
        },
        "has_email": true,
        "email_needs_agreement": false,
        "is_email_valid": true,
        "is_email_verified": true,
        "email": "hellp@kakao.com",
        "has_birthday": true,
        "birthday_needs_agreement": false,
        "birthday": "0312",
        "birthday_type": "SOLAR",
        "has_gender": true,
        "gender_needs_agreement": false,
        "gender": "male"
    }
}
*/
