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
				StudyPlaceId: uuid.MustParse("ae2bdeb1-a820-49c1-adca-4405be0034ee"),
				Permissions:  []models.Permission{models.PermissionAdmin},
			},
			lessonsAmount:        2250,
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
			from:       time.Date(2024, 1, 2, 0, 0, 0, 0, time.Local),
			to:         time.Now().AddDate(0, 1, 0),
			groups: []uuid.UUID{
				uuid.MustParse("5e83889a-961d-4655-830f-ad14b4cf8d61"),
				uuid.MustParse("bb140eac-40cd-4aa7-8c48-4297e734f7ec"),
				uuid.MustParse("8dea85c0-ffd9-4984-990e-5b4affc3f5ec"),
				uuid.MustParse("a14f182d-fb8d-47ce-ba60-030db91e293e"),
				uuid.MustParse("0ffc1802-3e58-458d-ad63-0054dcc22088"),
				uuid.MustParse("95b08054-1774-417f-ace4-c6e082812f27"),
				uuid.MustParse("447a5fae-1098-4325-9454-2d27d8a3e59b"),
				uuid.MustParse("20ca4cd8-6871-4ad0-b69c-b840add6a88b"),
				uuid.MustParse("7cf57a8a-20d8-4f7d-9d65-b2faeec229dd"),
				uuid.MustParse("916f4cca-8d5c-4400-873c-e3bb054c9def"),
			},
			rooms: []uuid.UUID{
				uuid.MustParse("721620c4-d741-442b-9c32-c0885bf5fba5"),
				uuid.MustParse("943f21be-b77f-4bb5-adcb-5f4f1cd34398"),
				uuid.MustParse("cf9f1100-fabf-49b3-93c6-688b62e1f74c"),
				uuid.MustParse("5dbe435b-07c2-4448-95e9-6b5f6f0c3668"),
				uuid.MustParse("04490a9f-ac0c-45df-88ab-054318ee2411"),
				uuid.MustParse("59f8b90c-b786-48c8-a5a3-5d88ee9679f9"),
				uuid.MustParse("17a1f74c-0ea3-42eb-9ddc-28618957cc57"),
				uuid.MustParse("4ca4c062-a3cb-4fd7-a04c-3325d0088eef"),
				uuid.MustParse("842bb39b-37e3-49e4-b4fa-7c33eeb1933a"),
				uuid.MustParse("9ce3be23-2aca-42d4-a39f-9546d266a820"),
				uuid.MustParse("c109ec81-412c-4839-8787-3d0cacc5dfc9"),
				uuid.MustParse("c1a32ee1-87ec-4633-b59c-289fa50d0188"),
				uuid.MustParse("9d431e1b-e78f-4c47-a41e-484956df464e"),
				uuid.MustParse("8f4b3396-5b7b-446a-920d-05fbe8a0be06"),
				uuid.MustParse("e4a3e79b-34c6-43a3-8e3d-37417e0b0c34"),
				uuid.MustParse("847af2cf-349c-47c9-beea-11b9920c1280"),
				uuid.MustParse("ab74422b-a856-43c5-ab92-9b3a8b6e3c92"),
				uuid.MustParse("c7d83edb-41ee-4994-865a-89b5fd785654"),
				uuid.MustParse("a96a4b5d-c0ce-4dd0-b217-19dc42fbeee3"),
				uuid.MustParse("4af8d383-cee1-497f-a77c-af82c63d6ad2"),
				uuid.MustParse("02eb3801-e7a6-4faa-807a-8a37ee71765e"),
				uuid.MustParse("abe31a02-fc62-4686-88b1-353392b8aaa5"),
				uuid.MustParse("bfe74287-723b-47fa-a365-3ea4fd7ba990"),
				uuid.MustParse("f392232e-90d8-419e-9e6d-3cada6650f40"),
				uuid.MustParse("0822f685-b8fa-45e6-9484-a2a907fd2571"),
				uuid.MustParse("bfb890bb-8502-4bff-a31b-e202a7c3e0ba"),
				uuid.MustParse("1a6f9404-95e6-4dc6-a0ef-aa1291fe1471"),
				uuid.MustParse("aa15bbe9-0113-4cb3-ae2b-91befa4816b6"),
				uuid.MustParse("479679b8-236d-4834-aad3-f339fc50dd0a"),
				uuid.MustParse("a5478d9e-ac2b-472a-98e1-1846046e0a32"),
				uuid.MustParse("740bfd85-e8f9-4897-8b5c-f19ca02ce628"),
				uuid.MustParse("2d9a4ece-4df9-4663-9014-c91e4e900622"),
				uuid.MustParse("ee2ff078-acda-49b5-8ae7-c4d6185aed94"),
				uuid.MustParse("56f0aedc-dcc8-4ba1-8dd1-34189c9e0c04"),
				uuid.MustParse("344c6d34-7823-4644-b572-a4e4aa1f8f5e"),
				uuid.MustParse("9c388f26-a14f-4d39-ba2d-71d38bae595d"),
				uuid.MustParse("dd97f85e-96d6-4d38-a8b3-d6c7eb673486"),
				uuid.MustParse("076f8207-371c-47cb-afe7-0bd24779d3de"),
				uuid.MustParse("b0e08d72-3b7b-43f1-a13c-0770e7e6b698"),
				uuid.MustParse("e40472ff-f89d-4ca9-b445-e9aefefa4307"),
				uuid.MustParse("019ee176-6cb9-466f-80f6-4879ec377811"),
				uuid.MustParse("cb26a313-b976-430b-9dff-ce50b4c2450a"),
				uuid.MustParse("dfeaf71e-1639-4b88-9f8a-8910dc53ea6b"),
				uuid.MustParse("e205aff9-ad5e-4e71-8f6d-0fa58acc3d65"),
				uuid.MustParse("234b7314-a802-4448-97a9-26e9d0f853e5"),
				uuid.MustParse("73fe4ebe-f779-4721-ba41-ff464c76a924"),
				uuid.MustParse("e9667bdf-183b-4405-82f7-f8ff352f9957"),
				uuid.MustParse("809af9cc-7983-4e78-afe7-1539b6536aed"),
				uuid.MustParse("02d02b8f-9252-457e-804b-d0d786019745"),
				uuid.MustParse("f1a7cf92-a679-46a2-ae3d-319d4ed7ab10"),
			},
			subjects: []uuid.UUID{
				uuid.MustParse("2d5df227-0f89-474d-8f20-47c200c95f41"),
				uuid.MustParse("0743fc4f-4897-40a7-bf7e-4cdd014a59ba"),
				uuid.MustParse("c7605db7-e7c6-41b2-bd7e-feebce8cfb66"),
				uuid.MustParse("f361037f-2fc7-480c-843d-977296ec3303"),
				uuid.MustParse("624300cf-2904-4cd9-8ddd-a0c99f6d19a3"),
				uuid.MustParse("65aba934-0b22-4962-9877-1c9a26fadc63"),
				uuid.MustParse("19b5f3f4-7518-4628-9119-dad0df3fadc1"),
				uuid.MustParse("d38ded5d-d2fa-4086-970b-8d170e5bd777"),
				uuid.MustParse("6054857c-1341-4887-bb85-bd204783e6b0"),
				uuid.MustParse("902ec51e-0947-437e-a329-73cba2a77ea8"),
				uuid.MustParse("bd3aa108-055d-4855-90fb-5e9ed68fc90c"),
				uuid.MustParse("25d34f62-fdc2-42fc-9178-5f401c2cd2a3"),
				uuid.MustParse("ae0dafed-4c16-4dbd-b972-4c10cddf0db3"),
				uuid.MustParse("adacccba-a284-476e-a9cb-ce7f563035c1"),
				uuid.MustParse("e273c80d-0c00-4e5d-8d69-9ddde28c6446"),
				uuid.MustParse("7a5276c5-6c3d-455e-9fd9-c54bef777a58"),
				uuid.MustParse("bc7940a1-abef-41c4-8de8-a3ad386c974c"),
				uuid.MustParse("aef7c1b8-0897-4c2f-9bd2-7ce3c0a3b8b8"),
				uuid.MustParse("6ffaa763-03a7-4509-a490-a2b8b94f3f7e"),
				uuid.MustParse("4aa03202-d197-44b1-bd66-0d55a4986c2b"),
				uuid.MustParse("b8e5d0e7-1686-46e3-9fb8-133d2d981bcd"),
				uuid.MustParse("98f57b10-acd6-4d6b-9c89-944203967ee1"),
				uuid.MustParse("26ef6ee6-f43e-40f6-b94b-51002dd24284"),
				uuid.MustParse("10cf8a3e-a56c-44be-9e19-94712cf962cd"),
				uuid.MustParse("db90743a-82cb-4b46-8864-3f15e1bf12ba"),
				uuid.MustParse("ad51fdf6-1cff-48a6-a4da-4f5106adcd39"),
				uuid.MustParse("a896d9b8-b264-409e-8d81-738cbaf5d300"),
				uuid.MustParse("9f863a64-c28c-42fa-b6cc-46e6039063a6"),
				uuid.MustParse("31b8c3d8-4a49-4434-ac41-e5310dd09c5f"),
				uuid.MustParse("844a2a6a-278d-49c3-8fe9-d2fe6f2ef712"),
				uuid.MustParse("9a1ec208-ccc6-412c-b27a-99c0ada8f151"),
				uuid.MustParse("831f3a72-e58a-47c4-9944-e48209256421"),
				uuid.MustParse("f60f0e00-46c4-421f-af2c-9a72f8b745de"),
				uuid.MustParse("f259137f-61a0-4f3b-87c9-ab82b968f024"),
				uuid.MustParse("0c31c9b5-d6d8-4858-ae01-b7bb6039b54c"),
				uuid.MustParse("2c15fa36-fa29-4434-bc04-53d892c06620"),
				uuid.MustParse("d5d979e8-ec5d-423b-9e88-fa094e9b4c1a"),
				uuid.MustParse("428ad7c0-bdfd-4a59-bde9-e5fc428489d2"),
				uuid.MustParse("714ca3b1-93db-44c7-9087-c3a26792857f"),
				uuid.MustParse("fbe0ee99-cd2f-4196-8525-9af51b53ef2c"),
			},
			teachers: []uuid.UUID{
				uuid.MustParse("920b3687-783f-42d9-a700-80a0cac49ec6"),
				uuid.MustParse("32d8da9c-9856-43be-adfc-45ed2dd08aec"),
				uuid.MustParse("58907af2-1c22-4ee5-8599-5dd9db955381"),
				uuid.MustParse("16b3ef4e-3353-4cbb-996d-00f123bb3db3"),
				uuid.MustParse("c3a9de7e-1031-46e3-bd61-ddc13cf57cad"),
				uuid.MustParse("5c0bd095-eac9-4291-a007-7448e128e629"),
				uuid.MustParse("2b6b863d-2bef-4d54-95d4-36a278487749"),
				uuid.MustParse("da042b3c-ef6c-4178-bdb3-47c2d66da3e7"),
				uuid.MustParse("e9393d87-4327-4799-96ca-35db9be1ad67"),
				uuid.MustParse("15a9438d-7cb7-4542-b944-053ebf7fc3d3"),
				uuid.MustParse("f36099a5-759e-4d28-ac12-8b52b4bd94d9"),
				uuid.MustParse("75f68f12-e762-433f-a491-7bc8f87f6046"),
				uuid.MustParse("e2040986-6b79-4f5e-92ff-60847899d154"),
				uuid.MustParse("71907b56-612c-4148-93be-f9ef91aac852"),
				uuid.MustParse("2dd24f21-7c3e-4c17-9fa1-7a8cdceed90a"),
				uuid.MustParse("5bbe2ed1-5ec9-4333-beb0-beae76be28d2"),
				uuid.MustParse("12d454ac-3d76-4a45-8eae-0ff483ce86cc"),
				uuid.MustParse("dd3d1ff2-612e-43d0-8839-dc17be9d0071"),
				uuid.MustParse("4aac4262-51e4-4244-921a-0a736aa7f02b"),
				uuid.MustParse("13242709-9786-49a6-9329-5314182184f9"),
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
