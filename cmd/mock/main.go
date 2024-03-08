package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-schedule/internal"
	"github.com/stdyum/api-schedule/internal/app/controllers"
	"github.com/stdyum/api-schedule/internal/app/dto"
)

type config struct {
	enrollment           models.Enrollment
	lessonsAmount        int
	generalLessonsAmount int
	lessonDuration       time.Duration
	times                []time.Duration
	dayIndexes           []int
	from                 time.Time
	to                   time.Time
	groups               []uuid.UUID
	rooms                []uuid.UUID
	subjects             []uuid.UUID
	teachers             []uuid.UUID
	primaryColor         string
	secondaryColor       string
}

func main() {
	_, ctrl, err := internal.Configure()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	//enrollments := []models.Enrollment{
	//	{
	//		StudyPlaceId: uuid.MustParse("7864c758-6405-41d0-b4f9-c467b78ad78e"),
	//		Permissions:  []models.Permission{models.PermissionAdmin},
	//	},
	//}

	configs := []config{
		{
			enrollment: models.Enrollment{
				StudyPlaceId: uuid.MustParse("f6815261-3205-4c78-af3e-096430384af2"),
				Permissions:  []models.Permission{models.PermissionAdmin},
			},
			lessonsAmount:        22500,
			generalLessonsAmount: 450,
			lessonDuration:       time.Minute * 45,
			times: []time.Duration{
				time.Hour * 8,
				time.Hour * 9,
				time.Hour * 10,
				time.Hour * 11,
				time.Hour * 12,
				time.Hour * 13,
				time.Hour * 14,
				time.Hour * 15,
				time.Hour * 16,
				time.Hour * 17,
			},
			dayIndexes: []int{1, 2, 3, 4, 5, 6},
			from:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			to:         time.Now().AddDate(0, 1, 0),
			groups: []uuid.UUID{
				uuid.MustParse("9bb5175d-978a-45b0-92cc-7661024c213a"),
				uuid.MustParse("b221cce7-5e5c-4533-a8fa-5a2d2e08483c"),
				uuid.MustParse("9c144179-3ce9-4f6c-b907-bb686453e66d"),
				uuid.MustParse("c80c087a-0b97-4d83-a7af-5dd925a2c0b7"),
				uuid.MustParse("26e1b92c-8834-4281-9d75-4337608e9089"),
				uuid.MustParse("fc211f9e-9e4a-4ec2-b4a2-268c11d22425"),
				uuid.MustParse("9790ac3b-57bc-4caf-bf25-9c1531168535"),
				uuid.MustParse("b85df517-cbca-4932-9a62-706759a1f6ca"),
				uuid.MustParse("a3b4f8e0-2eaf-425f-9c3b-09516c220704"),
				uuid.MustParse("6107aad8-d70b-4fa4-bb8a-c371fac9df19"),
			},
			rooms: []uuid.UUID{
				uuid.MustParse("86a3a8cc-b5a4-4e27-b2fb-f8b994ca287a"),
				uuid.MustParse("5ab12518-4434-4ec8-af52-8e3cb821dab3"),
				uuid.MustParse("3c8ce2b5-2847-4da2-90d4-8bac61ecbb3d"),
				uuid.MustParse("1bd8c750-0fb1-4b90-95d5-f8601edaf517"),
				uuid.MustParse("5e256a4a-30aa-4374-9a69-74739c3c2842"),
				uuid.MustParse("e9d066c8-0af5-4c1a-ae26-949c27f41ac1"),
				uuid.MustParse("304f0c1f-d109-4802-a129-f9a7529f45f6"),
				uuid.MustParse("17669cbe-fd65-43a4-901f-e8cce01f39b9"),
				uuid.MustParse("42860081-39ce-4da4-9b18-93b59f78d3d9"),
				uuid.MustParse("22de9d24-e438-492b-817a-e62a0f9c4d4a"),
				uuid.MustParse("7b8d95d9-114c-4442-8aed-bc9a232e83c8"),
				uuid.MustParse("3f9fc7c5-c43c-47d0-8da3-0bed0a1eaf00"),
				uuid.MustParse("74009565-9f75-41c1-890b-478fc367dd53"),
				uuid.MustParse("1edb05ee-fb4a-46d1-89ec-e388f183fffa"),
				uuid.MustParse("d2a915a7-9f6a-4909-811d-38113ebe0ed6"),
				uuid.MustParse("4c9733a6-7308-44ca-9af7-7814a079cb64"),
				uuid.MustParse("f90afaad-f4ab-4188-a908-fc5260faacd8"),
				uuid.MustParse("46f31ae4-19d8-4a89-86bb-db57638184d6"),
				uuid.MustParse("78993e9d-1e4a-4eb0-9b7b-cb9e57d081c6"),
				uuid.MustParse("73c7779a-f19b-4af7-b363-72e16c298750"),
				uuid.MustParse("a7edbb9d-8592-4e2c-9b5e-63cf775fe7b8"),
				uuid.MustParse("011d8605-5b57-4774-9d1b-6a1c06949964"),
				uuid.MustParse("50f83eff-3d33-4e8c-b605-959c9ff50031"),
				uuid.MustParse("fd825c47-8e1f-45ed-9c87-2f70805cef45"),
				uuid.MustParse("0538d133-3acb-44a5-8b9e-fbf2a2ee435c"),
				uuid.MustParse("74b2bb51-838d-4437-ac82-fad9856d8128"),
				uuid.MustParse("bba4d6cb-7d3d-4dcf-b273-f6e6882ab5c7"),
				uuid.MustParse("321590e2-d54c-4bb0-b332-e0b26620512c"),
				uuid.MustParse("3a45b941-323c-447c-a8ec-c7af877e6ddd"),
				uuid.MustParse("03332bd5-b7ef-4ccd-a599-237bd7efdc47"),
				uuid.MustParse("a65d071c-ae4a-4c5d-9614-660f46087f25"),
				uuid.MustParse("86c8767d-9678-4bf7-82e8-cf99bcf18ab5"),
				uuid.MustParse("272e4b30-3166-4c3a-b916-7238b033d2fb"),
				uuid.MustParse("0789cc06-5895-4e08-9d85-cd6c39a6a349"),
				uuid.MustParse("ffcc2798-d13b-4a41-861e-03d24ba78b70"),
				uuid.MustParse("ee6e900c-3ab0-4a6e-9e41-3d25fb03e7c5"),
				uuid.MustParse("14de692b-e023-4b7c-a193-0a3eaf5ea9e7"),
				uuid.MustParse("784550d6-ff57-4574-9740-0ecef1e1114c"),
				uuid.MustParse("39c0e95a-2604-4b0a-b9c6-aee2919cea02"),
				uuid.MustParse("6ded5d51-8fb0-4a1c-ba5d-bd8826ae3507"),
				uuid.MustParse("0c629402-f4ba-4e2e-82f0-7c0d92ed5133"),
				uuid.MustParse("2fa04a37-a4b1-491d-b3e2-abc761c15987"),
				uuid.MustParse("2ddeb40e-afc6-4344-83dc-06136b936dc1"),
				uuid.MustParse("0a402e40-e9bc-4b69-a644-7c09e9fc5bc2"),
				uuid.MustParse("db1b791d-f280-498c-915c-7694abe323e5"),
				uuid.MustParse("5594de39-2376-4f5c-9dd5-74639c1d92e1"),
				uuid.MustParse("84508774-364e-429f-84aa-d47cea08d893"),
				uuid.MustParse("f6ac5c78-b273-4179-8de7-376ae7aa4417"),
				uuid.MustParse("dd227d3c-e0a0-412a-abe3-813b3254644e"),
				uuid.MustParse("9b3a3381-eafc-4f4a-a658-0ac35eedf1c3"),
			},
			subjects: []uuid.UUID{
				uuid.MustParse("24d7cd25-ce0a-4d0d-abed-cf3e31730059"),
				uuid.MustParse("6e664376-9786-4778-b9fe-532d8e01cb9d"),
				uuid.MustParse("9f965713-440a-4177-a8ee-f51aceb3be52"),
				uuid.MustParse("653d9ea7-43fe-4364-8d8d-194935a0999a"),
				uuid.MustParse("7aa416f3-dd0a-4a51-aa5b-738e2e955588"),
				uuid.MustParse("316b8f6e-cbec-4084-8b4e-c9cdbe76fbe9"),
				uuid.MustParse("55a74197-305c-4491-ad4d-5e06874d19c5"),
				uuid.MustParse("f46bba6b-58e0-410d-9737-dc47d7dda4af"),
				uuid.MustParse("c7d90b95-374e-4624-9e7b-b8b8fd4b91be"),
				uuid.MustParse("ddadad8b-7848-40f0-8d2d-a7da89710af5"),
				uuid.MustParse("38ce341b-00ae-49bf-88b6-1ae332cb85cd"),
				uuid.MustParse("7c8b1a06-87ea-4d45-8cbe-bbf56e8d9afa"),
				uuid.MustParse("ec31f626-79e6-4dd3-ba0f-1d9108c5ed6a"),
				uuid.MustParse("db723716-5abf-44fb-a899-2f471bd39bbc"),
				uuid.MustParse("3b97b57d-8cd5-41fd-a0f8-d3e822e3bfe5"),
				uuid.MustParse("6ed7934c-ae56-4bf3-9c2b-7f58ec8b2205"),
				uuid.MustParse("44653f67-b4b8-4d42-8eb3-604b03c7e5e2"),
				uuid.MustParse("ac50c54b-bfba-45b6-a668-1397f4d1e67f"),
				uuid.MustParse("f534ba0b-5d66-40ea-bff2-2c2dd3493ee6"),
				uuid.MustParse("f0e14b61-8155-4253-b9cc-d8d4854ad5de"),
				uuid.MustParse("5639b364-59f8-4e02-b504-33b5053ecd12"),
				uuid.MustParse("bf40cff9-dca0-46da-a8e3-b9172267d97a"),
				uuid.MustParse("43701e48-5c92-4bb0-b124-a6bd86d3945b"),
				uuid.MustParse("075332cf-5c22-4147-a70e-df8d8bf2523d"),
				uuid.MustParse("274626f7-036b-4713-8b97-2bf32d477d5d"),
				uuid.MustParse("a7013d68-902f-4e23-893f-e97a0dafe958"),
				uuid.MustParse("f9b8eb4a-c93a-4e65-803f-db2bc2b93c6d"),
				uuid.MustParse("338d50e0-6d91-4887-b0dc-daaecd168f43"),
				uuid.MustParse("2f99fb69-60c6-4efe-b191-1bfb43a0a4f5"),
				uuid.MustParse("207da103-ec6a-4c02-ad6d-190999feeafd"),
				uuid.MustParse("8d270364-cd0e-4deb-9dd8-e6f301d6ba46"),
				uuid.MustParse("81576e4d-c974-45da-96c2-c4303d740147"),
				uuid.MustParse("e88b43ee-cc35-42ec-9ea5-6cc7adc6c713"),
				uuid.MustParse("590b0ebc-3fd7-411b-a26d-4009a01462f0"),
				uuid.MustParse("691b3d46-5b63-4f57-8bc1-7aaa27ebd229"),
				uuid.MustParse("31aa6112-1ab2-4adf-a76f-b19e1595e494"),
				uuid.MustParse("c9042c7e-258a-4b96-a9a7-36950b93e334"),
				uuid.MustParse("11cf76f3-6b43-4ddb-b423-654634775509"),
				uuid.MustParse("c576653c-8026-467f-882c-f4fd7558c0b5"),
				uuid.MustParse("497fa8be-ba4d-4e59-a2fe-bc5202622b3f"),
			},
			teachers: []uuid.UUID{
				uuid.MustParse("cfc3ab3b-0bd9-4b4f-bbc0-bf1482736b4e"),
				uuid.MustParse("c7285f3c-c32b-4a83-b979-2bf3165ce784"),
				uuid.MustParse("b7f7cd05-b530-4572-ad9f-04a8db1e3354"),
				uuid.MustParse("a7069240-bc5b-461b-ba56-b7c47a2d5650"),
				uuid.MustParse("88d28ca5-9f36-431b-a283-78fbf07db712"),
				uuid.MustParse("9e33d7cc-71e7-465f-99eb-2a5d1c77a7eb"),
				uuid.MustParse("6470c3bf-2f3c-49c6-a452-797691f831f5"),
				uuid.MustParse("5feb3000-0f29-4ae6-939a-a9eb3ba281f4"),
				uuid.MustParse("a24432b4-f07d-453e-99b6-c457a419e73e"),
				uuid.MustParse("61a709c8-d1d1-4061-abe7-b7fd64879fe0"),
				uuid.MustParse("4d1faa20-2ed2-4140-8bb1-430cd05b63b6"),
				uuid.MustParse("e161b069-50b1-42a1-9b25-e4857502f023"),
				uuid.MustParse("375bcd56-4bdc-467c-9ca2-d98710d43b17"),
				uuid.MustParse("06f59bad-759f-4163-a5d2-80ddc0c6424f"),
				uuid.MustParse("260af6ca-93ed-4a01-9061-2d28313fb528"),
				uuid.MustParse("36cb3144-0d68-4748-905e-e152a8970fec"),
				uuid.MustParse("f5fe521a-85eb-4d62-97d5-6383bcb0b195"),
				uuid.MustParse("81e65661-935b-4ac2-a226-8da3dc9d01ac"),
				uuid.MustParse("ff062039-5513-4ac8-9f4f-50bbb3a273cf"),
				uuid.MustParse("29d32e4a-e1e7-4834-a424-67ba0abec64b"),
			},
			primaryColor:   "white",
			secondaryColor: "transparent",
		},
	}

	for _, c := range configs {
		mock(ctx, ctrl, c)
	}
}

func mock(ctx context.Context, ctrl controllers.Controller, cfg config) {
	log.Println("starting mock for ", cfg)
	generalLessons := make([]dto.CreateLessonGeneralEntryRequestDTO, cfg.generalLessonsAmount)
	for i := 0; i < cfg.generalLessonsAmount; i++ {
		groupIndex := rand.Intn(len(cfg.groups))
		roomIndex := rand.Intn(len(cfg.rooms))
		subjectIndex := rand.Intn(len(cfg.subjects))
		teacherIndex := rand.Intn(len(cfg.teachers))

		startTimeIndex := rand.Intn(len(cfg.times))
		dayIndex := rand.Intn(len(cfg.dayIndexes))

		lesson := dto.CreateLessonGeneralEntryRequestDTO{
			GroupId:        cfg.groups[groupIndex],
			RoomId:         cfg.rooms[roomIndex],
			SubjectId:      cfg.subjects[subjectIndex],
			TeacherId:      cfg.teachers[teacherIndex],
			StartTime:      cfg.times[startTimeIndex],
			EndTime:        cfg.times[startTimeIndex] + cfg.lessonDuration,
			DayIndex:       cfg.dayIndexes[dayIndex],
			LessonIndex:    startTimeIndex,
			PrimaryColor:   cfg.primaryColor,
			SecondaryColor: cfg.secondaryColor,
		}

		generalLessons[i] = lesson
	}

	dayAmount := int(cfg.to.Sub(cfg.from).Hours() / 24)
	schedule := make([]dto.CreateScheduleMetaEntryRequestDTO, dayAmount)
	for i := 0; i < dayAmount; i++ {
		meta := dto.CreateScheduleMetaEntryRequestDTO{
			Date:   cfg.from.AddDate(0, 0, i),
			Status: "updated",
		}
		schedule[i] = meta
	}

	currentLessons := make([]dto.CreateLessonEntryRequestDTO, cfg.lessonsAmount)
	for i := 0; i < cfg.lessonsAmount; i++ {
		groupIndex := rand.Intn(len(cfg.groups))
		roomIndex := rand.Intn(len(cfg.rooms))
		subjectIndex := rand.Intn(len(cfg.subjects))
		teacherIndex := rand.Intn(len(cfg.teachers))

		startTimeIndex := rand.Intn(len(cfg.times))

		date := schedule[rand.Intn(len(schedule))].Date

		lesson := dto.CreateLessonEntryRequestDTO{
			GroupId:        cfg.groups[groupIndex],
			RoomId:         cfg.rooms[roomIndex],
			SubjectId:      cfg.subjects[subjectIndex],
			TeacherId:      cfg.teachers[teacherIndex],
			StartTime:      date.Add(cfg.times[startTimeIndex]),
			EndTime:        date.Add(cfg.times[startTimeIndex] + cfg.lessonDuration),
			LessonIndex:    startTimeIndex,
			PrimaryColor:   cfg.primaryColor,
			SecondaryColor: cfg.secondaryColor,
		}

		currentLessons[i] = lesson
	}

	log.Println("mock data generated, inserting")

	uslices.ChunkFunc(schedule, 50, func(chunk []dto.CreateScheduleMetaEntryRequestDTO, i int) {
		log.Println("meta[" + strconv.Itoa(i) + "]")
		_, err := ctrl.CreateScheduleMeta(ctx, cfg.enrollment, dto.CreateScheduleMetaRequestDTO{List: chunk})
		if err != nil {
			log.Println("meta["+strconv.Itoa(i)+"]; error ", err)
		}
	})

	uslices.ChunkFunc(generalLessons, 50, func(chunk []dto.CreateLessonGeneralEntryRequestDTO, i int) {
		log.Println("general lessons[" + strconv.Itoa(i) + "]")
		_, err := ctrl.CreateLessonsGeneral(ctx, cfg.enrollment, dto.CreateLessonsGeneralRequestDTO{List: chunk})
		if err != nil {
			log.Println("general lessons["+strconv.Itoa(i)+"]; error ", err)
		}
	})

	uslices.ChunkFunc(currentLessons, 50, func(chunk []dto.CreateLessonEntryRequestDTO, i int) {
		log.Println("current lessons[" + strconv.Itoa(i) + "]")
		_, err := ctrl.CreateLessons(ctx, cfg.enrollment, dto.CreateLessonsRequestDTO{List: chunk})
		if err != nil {
			log.Println("current lessons["+strconv.Itoa(i)+"]; error ", err)
		}
	})
}
