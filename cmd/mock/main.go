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
				StudyPlaceId: uuid.MustParse("fbd0abd6-7fda-485f-9e3a-79d41cd641ba"),
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
			from:       time.Date(2024, 1, 2, 0, 0, 0, 0, time.Local),
			to:         time.Now().AddDate(0, 1, 0),
			groups: []uuid.UUID{
				uuid.MustParse("05317a7e-fe81-4fab-af0f-d378beed0de5"),
				uuid.MustParse("5039a6da-a015-4d07-bf6f-073a23d5efd1"),
				uuid.MustParse("0e582adf-37bb-4202-b9cb-1e4653bc6e37"),
				uuid.MustParse("d1468a77-ffc3-4a10-8ebd-b28f2b0756ae"),
				uuid.MustParse("e6efa411-8c05-4ca7-8ab4-dcc367bda64d"),
				uuid.MustParse("f509a972-d221-4c73-a696-9f76114cdfa2"),
				uuid.MustParse("8bc27feb-4a47-47cb-aea8-4b26d9d0095b"),
				uuid.MustParse("f80b8a6b-1e00-41db-b940-59669befeb6f"),
				uuid.MustParse("44cdeb25-a3d5-476b-8f82-6d35563bd6ce"),
				uuid.MustParse("9fb44b63-155f-4e83-bc2c-4224990361b1"),
			},
			rooms: []uuid.UUID{
				uuid.MustParse("0630f370-b635-4f4d-86a5-88c15df6b961"),
				uuid.MustParse("e4d68ba7-8bfe-4d77-a7be-1c3744ec4ca2"),
				uuid.MustParse("7abf3d9a-dfc8-4e2b-ab93-23664d4e627c"),
				uuid.MustParse("47a3034a-a230-4591-affd-8b886e09e080"),
				uuid.MustParse("7ef6c289-4795-4b8d-a5de-b765382a3a26"),
				uuid.MustParse("016f7bc5-84a5-445d-be64-28e3e782306c"),
				uuid.MustParse("eb77b7d0-5fc1-49c9-80aa-1d4fb412076d"),
				uuid.MustParse("a1d29859-db63-4240-9a85-49fdcde29c82"),
				uuid.MustParse("96b28d38-882c-4551-967c-39e441682a8b"),
				uuid.MustParse("a7e97faa-2119-4f87-be57-b82d78ba4ae5"),
				uuid.MustParse("19907fd1-5c58-457e-816e-08fe4fbc8b97"),
				uuid.MustParse("84ed0165-083b-42da-b6cf-92500b57b7be"),
				uuid.MustParse("fda86034-77c0-4ca9-b247-587bb4518b8f"),
				uuid.MustParse("23456fc0-1305-4dac-9186-7061b444b682"),
				uuid.MustParse("978178fc-6fd9-4387-a053-5195de2b45c8"),
				uuid.MustParse("a78108c3-5b2d-4afe-9f87-94e47ae64fd1"),
				uuid.MustParse("8e9d5d1d-e436-41ca-ab90-ffe898fff635"),
				uuid.MustParse("8c24ad47-5fa2-46ce-8245-61992ce55f95"),
				uuid.MustParse("3e5be580-99a5-474e-8dba-5df57f4acd5c"),
				uuid.MustParse("e6a66a7b-59ac-41bd-a452-45b1a15264e7"),
				uuid.MustParse("a89120d2-bd00-468c-93e6-ac7343d00bd3"),
				uuid.MustParse("646bd94a-72df-4b4f-b42b-d2810459d70e"),
				uuid.MustParse("2a3fa86e-fcde-4282-b0d5-eba401a1ebd2"),
				uuid.MustParse("00f6caf6-7836-4be4-b982-1eb31e6e34ef"),
				uuid.MustParse("1b950912-5420-416b-8972-26fdcf40f130"),
				uuid.MustParse("bcdfcb95-1c41-4354-8400-ebc5df1dfa3c"),
				uuid.MustParse("d254c63a-3e2f-4582-b1ab-f712e562ceaa"),
				uuid.MustParse("be5c97f2-3044-4bbf-b3bc-8738d0ea4316"),
				uuid.MustParse("89ea1e64-9c8c-48f5-9d19-f08ea7c19c30"),
				uuid.MustParse("4bb66700-3f0c-451a-9e28-ad213843f298"),
				uuid.MustParse("11987196-3f21-4d5d-b62d-c218e1822144"),
				uuid.MustParse("e826972a-6b68-482a-9f30-e01cb5704be5"),
				uuid.MustParse("9d308ba4-7fa4-4be5-9ce8-a4d93d799377"),
				uuid.MustParse("c4a324b7-72c1-442b-b0f0-f58b472e7e54"),
				uuid.MustParse("ea7cc3c0-d46d-4091-9950-a1ca561c2de5"),
				uuid.MustParse("8bb6f1c4-20b6-4fb6-ad5d-f988fc185131"),
				uuid.MustParse("a81907bc-3819-4563-92a6-6032253c1688"),
				uuid.MustParse("bd36b837-3bb7-4149-a168-fc9a507c5b79"),
				uuid.MustParse("e8f0d5f8-bc6f-4772-836c-b1db09785052"),
				uuid.MustParse("db9d202f-6ed9-4467-b70a-4637c37763c8"),
				uuid.MustParse("aee156f4-eb50-44e7-b52f-d67302498c7c"),
				uuid.MustParse("dcedc17d-1866-4f89-98be-16e65eaaba11"),
				uuid.MustParse("d4ad9b73-de23-43cd-ab99-4fab98beeff6"),
				uuid.MustParse("58cb1585-7114-45c9-85b2-d18edc2bfb50"),
				uuid.MustParse("555ff546-d15a-4a17-a5b7-975193492b72"),
				uuid.MustParse("fef7295f-67cb-4a0c-a543-a93accdd4554"),
				uuid.MustParse("e8879e25-088c-4b40-b97d-1717737e3b98"),
				uuid.MustParse("0df46385-161f-4b93-be6d-570fcc073cff"),
				uuid.MustParse("b81253a1-5508-4f61-9970-0e9deb66c7c2"),
				uuid.MustParse("ff3d3bc0-ae98-430a-a340-46731d632439"),
			},
			subjects: []uuid.UUID{
				uuid.MustParse("b5683af8-cbd5-4774-9ab5-1d6df6bd61bb"),
				uuid.MustParse("830f1d4f-0c45-43af-8626-aec0d55272b8"),
				uuid.MustParse("507f08f1-a41c-4c88-87cd-01e5b8b34f8d"),
				uuid.MustParse("e1c241d4-a2fc-465c-a9b9-6166280b84a7"),
				uuid.MustParse("4104f9fa-ea6b-43da-870d-cf4d53f55e07"),
				uuid.MustParse("98af2c22-6cd4-4d29-a05f-8cf52319a610"),
				uuid.MustParse("13fe5ea9-1817-46aa-8b30-bdb21e458b23"),
				uuid.MustParse("788ea2ac-3200-48c0-8a33-e3673646a6dd"),
				uuid.MustParse("56a02da8-2a7a-48d6-a13d-a969275e545e"),
				uuid.MustParse("b59044b4-d331-411e-8c3a-fdb7f5f0203a"),
				uuid.MustParse("cf0647da-f98a-4a55-a76c-32aafe43ff26"),
				uuid.MustParse("c9e7dde4-78fb-4339-8648-4a880a52a2e5"),
				uuid.MustParse("8c1f6c1f-9d73-48ee-87fa-d369ee962188"),
				uuid.MustParse("770e0c21-9227-428f-8f8b-e2edae21f03c"),
				uuid.MustParse("b8989e58-4e64-4350-a905-156e024de874"),
				uuid.MustParse("5dd10951-6860-4f20-84c9-30f730d1c92a"),
				uuid.MustParse("d403cbfd-f099-480b-b3d8-5690203deba0"),
				uuid.MustParse("024d80ec-d694-489a-a275-068a59a32e73"),
				uuid.MustParse("a4dbe8bc-7b8a-4a29-b184-14c717b6961c"),
				uuid.MustParse("bd5dbf2f-3291-4ab2-a267-7b71d8e6101d"),
				uuid.MustParse("d27f0a7d-d405-4058-a8d2-a704c9d72bf9"),
				uuid.MustParse("1ecc4466-c132-427f-a6db-ae9c8290bc41"),
				uuid.MustParse("13957d50-1d3c-4155-b62f-0485f21d7970"),
				uuid.MustParse("0035431a-07ac-4ece-b953-e666d2993f21"),
				uuid.MustParse("04a42ec9-b29f-487d-ae70-6c1828fd840a"),
				uuid.MustParse("a74bb844-a855-4ce4-840d-ac312c5c11dc"),
				uuid.MustParse("15186f30-46b7-47e9-9230-0ec8447e672e"),
				uuid.MustParse("c11f86fc-8f8b-4cbe-ae96-264b2a64e7a9"),
				uuid.MustParse("d7ff6641-88ad-44bc-906e-6ebbfca405c9"),
				uuid.MustParse("288b414b-1669-4935-ac9f-ffaee8816fe6"),
				uuid.MustParse("275c44ac-00f0-47b0-909a-f1d85025acc3"),
				uuid.MustParse("4a8883cd-3e1f-4a8e-8281-536c8f4eb55b"),
				uuid.MustParse("46751e8b-5174-4116-8c57-235b88df717a"),
				uuid.MustParse("5cdcc0da-3cd5-418c-aca0-e6bbba7d954d"),
				uuid.MustParse("81b78f35-6b5e-4e1a-9627-f96231c416cc"),
				uuid.MustParse("b8fa1e3f-ca61-4105-bbcf-b31f286b89e8"),
				uuid.MustParse("44d69c51-794f-4d8a-96af-842e98c495b0"),
				uuid.MustParse("6f26c9f5-fc44-47e4-9dac-1246d6fb8e01"),
				uuid.MustParse("f5157b43-a955-4640-b2e3-65a8ce820ff1"),
				uuid.MustParse("130e5d75-9a8b-440c-b910-f48d58f09827"),
			},
			teachers: []uuid.UUID{
				uuid.MustParse("dccdf28b-2916-42d1-971d-552763ab68fe"),
				uuid.MustParse("bbae361a-da78-4ded-a454-ed1c427ca83f"),
				uuid.MustParse("1e3872ac-0889-4cdc-9fd1-d52dc34c4727"),
				uuid.MustParse("b2081b36-67af-4a11-abdb-76ab4dc62282"),
				uuid.MustParse("29f7e84a-7675-4679-9b11-7863a4d025db"),
				uuid.MustParse("cd07f3d4-b527-4f7e-a0ff-4f07b8be5985"),
				uuid.MustParse("4ef0b61b-a4c0-4c54-8201-2688bb3fe1b8"),
				uuid.MustParse("0c37987a-8a6b-4ffe-9603-e0647705da64"),
				uuid.MustParse("e313d441-7a1a-414f-a100-caf9649cd8e8"),
				uuid.MustParse("eed0dbe9-bc19-4f63-8980-350d688aa540"),
				uuid.MustParse("a9f79b00-ef0a-4560-8425-da485a5d4f96"),
				uuid.MustParse("88fc097c-2c08-4b6d-9470-7a5b84393068"),
				uuid.MustParse("cee2f963-4f72-4a1a-8ef4-44276e738432"),
				uuid.MustParse("4d1ff263-46cc-4207-a421-064b50caa35f"),
				uuid.MustParse("0393d69b-7d8d-49e8-9614-06035be6aa1a"),
				uuid.MustParse("9ae56d8a-04bd-4a5f-9650-c60483da3e6d"),
				uuid.MustParse("64049ba8-f349-4ad1-b291-766df42b6bac"),
				uuid.MustParse("a8345c1e-8d89-4bb2-89e7-2137010a7449"),
				uuid.MustParse("8ad3c4d2-3eea-4f34-951d-2b7e0b71a155"),
				uuid.MustParse("d3088b7e-96bb-4362-8e0d-983ff3fbba90"),
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
