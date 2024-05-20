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
				StudyPlaceId: uuid.MustParse("7115bc03-143f-4acd-912d-02ffa9f14462"),
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
				uuid.MustParse("6648ada5-b87c-450c-93b2-9773a7b0290c"),
				uuid.MustParse("8e4808f2-0ff9-4c74-ab70-2a4e046c66a5"),
				uuid.MustParse("83e47c50-8105-4ad4-987d-b253ad1aaa5f"),
				uuid.MustParse("6fe101da-c349-4547-92eb-39ae8e2e8b4c"),
				uuid.MustParse("7b0b4e4f-e61e-4742-a652-ea5ffbefedeb"),
				uuid.MustParse("cb5608a5-3da0-40d1-9d96-dea76f122837"),
				uuid.MustParse("03372998-4dd6-4976-a56e-12ded08da7cb"),
				uuid.MustParse("0a8c538c-c568-4069-afa6-8b6074455b96"),
				uuid.MustParse("a2ebd778-9eb1-472f-a54b-e62267745053"),
				uuid.MustParse("00039363-5631-4358-9b3c-47c3baf2b1ff"),
			},
			rooms: []uuid.UUID{
				uuid.MustParse("bec8507d-60f9-4a18-96a2-4d3b68b1987f"),
				uuid.MustParse("9cf9fbaf-51e2-4470-850c-23f0818fe6c6"),
				uuid.MustParse("3e20266b-ae60-4b6f-bf4f-444ad368e91e"),
				uuid.MustParse("757ba697-9735-4c38-b974-a1d54ad1a61d"),
				uuid.MustParse("076e229f-bcc7-4990-b862-cee23844e270"),
				uuid.MustParse("123cf672-eaf7-4d44-abff-f2281bddf5ec"),
				uuid.MustParse("362f843b-205e-4fd2-87dd-0836bb4fac72"),
				uuid.MustParse("35980cb1-8844-4b40-9156-cde7a1aeba5a"),
				uuid.MustParse("be7910c1-d2d5-42c7-912d-b63e18d3fef5"),
				uuid.MustParse("11942d16-a5a8-45ff-b06c-ade521231609"),
				uuid.MustParse("06c0c16a-f5ee-44d4-b831-f5372792a459"),
				uuid.MustParse("ab55f353-e573-4ba5-8cc0-1e1a1e560534"),
				uuid.MustParse("3711431f-7bb9-4163-8b7a-51c2e3ada334"),
				uuid.MustParse("3e517f7b-7c89-4ab3-8437-698680d36697"),
				uuid.MustParse("806360af-2799-47ca-a61e-1a6044ac85b9"),
				uuid.MustParse("98c910be-e847-4725-aa68-e55ec5654d8f"),
				uuid.MustParse("76b68852-6fb3-4714-9998-36ddf140e317"),
				uuid.MustParse("69e04d4c-dfc9-4cb4-a2f7-4673cb54e911"),
				uuid.MustParse("20e7a954-649f-4287-93fc-d0115992df2d"),
				uuid.MustParse("17c9fd4f-8226-4eee-8072-841fca105c74"),
				uuid.MustParse("624fa9c2-2fb9-44ae-bbe6-49961b647a05"),
				uuid.MustParse("9e983b3d-7db2-4d26-9cce-c0698054904a"),
				uuid.MustParse("8a140acd-5f35-4023-b71f-b66ab7d6ee77"),
				uuid.MustParse("7c41e01c-5b36-440b-b8bc-4debd7c73fd1"),
				uuid.MustParse("7982a84c-14a1-4509-8b25-8cf7c496d615"),
				uuid.MustParse("226e974a-233b-4069-b8b5-889b1fb0225e"),
				uuid.MustParse("ea2062eb-6257-4d8e-aa1a-8c92954b2e8f"),
				uuid.MustParse("48c0f325-e031-44a3-9607-2c34b38c99c2"),
				uuid.MustParse("c6962e36-3b3e-42f1-bab5-053585fe1c8c"),
				uuid.MustParse("23ad5e92-977c-4765-9fb6-7f4ec3e4e59f"),
				uuid.MustParse("f9c73a9d-8414-4ea0-ac99-390e06681524"),
				uuid.MustParse("9e587629-9571-4eb8-9424-68de51e70ea1"),
				uuid.MustParse("4f3943e7-6c91-466d-a91d-86a435c27cb2"),
				uuid.MustParse("017200b6-4abb-468f-a21e-b9b0a2e1b2b8"),
				uuid.MustParse("d6fad6ef-f24f-4591-80bd-bef79589ffc5"),
				uuid.MustParse("2f4f2c95-df8f-43f2-9167-5d567afdce2b"),
				uuid.MustParse("26299af3-0f80-4c12-88e3-fe31dff9a5aa"),
				uuid.MustParse("f5747399-4d2d-4b4f-8124-3504e354a52a"),
				uuid.MustParse("9fad20e0-e21f-4d68-96fa-fdaec44722e4"),
				uuid.MustParse("58ef7dcf-288e-4757-a91d-dd55085c7848"),
				uuid.MustParse("b36b70f8-a0f2-41d0-8495-a8bd553f9dbb"),
				uuid.MustParse("dbe55db4-1df3-491c-bf4a-eae1d8931f89"),
				uuid.MustParse("989785e0-a810-43b8-af49-937dfa56fb7c"),
				uuid.MustParse("a1745416-f8d3-42d1-89ad-9d526710cf85"),
				uuid.MustParse("29be4a4c-de47-4051-85b4-2f27e5fbf6a0"),
				uuid.MustParse("96da67b5-7a2a-447d-a571-c4cdb7b6b774"),
				uuid.MustParse("b06442a3-4be7-48b6-bb52-6e9846b78039"),
				uuid.MustParse("7043a523-9f60-43f8-93c3-3da31e0bd778"),
				uuid.MustParse("8c32ea15-d8cc-4f9a-b631-05986c9d9092"),
				uuid.MustParse("0f4493ee-15a8-4e58-bbdf-718f4d4b644f"),
			},
			subjects: []uuid.UUID{
				uuid.MustParse("f79cec19-088b-4abd-bba3-f26bdd2e44b1"),
				uuid.MustParse("73f40c07-7592-41f2-b6b7-3c1eb7be983c"),
				uuid.MustParse("56f5815f-2814-4b39-9110-78ad647b616b"),
				uuid.MustParse("c54f5070-3ceb-4614-b505-ea002c9cbbf5"),
				uuid.MustParse("538dd287-e530-4140-a136-03c8cefcd424"),
				uuid.MustParse("161cf91c-65e8-4626-bb34-aecd0ff63fed"),
				uuid.MustParse("11dd68bb-4a2e-449c-8b2b-697241f92b7a"),
				uuid.MustParse("cd09e8f3-40af-4588-b9c7-83a86a4c3e75"),
				uuid.MustParse("bd00139d-f436-4206-88c6-c8fb5a4a3c16"),
				uuid.MustParse("f99a96c2-3fa9-4ad0-8c7f-7a137d495f0b"),
				uuid.MustParse("b4e156e5-26df-4cfd-a421-2e96c53f23f1"),
				uuid.MustParse("df21260e-4f94-413f-b80c-e7b35acdda92"),
				uuid.MustParse("0570e79f-411d-4a0a-89b5-e0e34f6bc15c"),
				uuid.MustParse("8f96e352-239d-40e1-b1d3-b72628b3233f"),
				uuid.MustParse("18aa3e3f-aa98-4182-8289-a65d4e391e58"),
				uuid.MustParse("3ef8f407-e622-4d5c-ac44-f6f0daf07b54"),
				uuid.MustParse("aff961fe-7f2c-4d31-bb50-76c7395025db"),
				uuid.MustParse("8ab6f93a-de5b-48f4-b76f-2531c62ff5a3"),
				uuid.MustParse("3f297bb4-3df9-46ce-af7f-ab07ae659f71"),
				uuid.MustParse("0d20ffb1-e022-4efa-bae2-3beb2c1eb8ba"),
				uuid.MustParse("d4a1ce34-d635-4741-b46b-c2e6067a7313"),
				uuid.MustParse("7c6aba1c-7095-4eba-9c29-87d233e7506b"),
				uuid.MustParse("5b0f5dbe-b1c9-4d23-b82c-bddc482fd79c"),
				uuid.MustParse("6a2d2bf2-5f44-4766-833b-f59006d607cf"),
				uuid.MustParse("381e52c6-cfc1-46ae-891b-03dc4d527b98"),
				uuid.MustParse("496922b3-5ed6-4328-be8c-d7a1995f1617"),
				uuid.MustParse("5c504b64-81c1-4201-b7c4-cd145bcea749"),
				uuid.MustParse("c9f701ee-853a-4438-9e9d-baf08e5ac7e9"),
				uuid.MustParse("92c94d29-82ba-4fc5-a67c-0488a53a88a4"),
				uuid.MustParse("20d26c75-6631-4ee8-99ae-76cd73a8b2bc"),
				uuid.MustParse("c98e983c-9327-4fe2-9def-1afa56d3d3a1"),
				uuid.MustParse("5834cf46-4546-426b-9570-adc9738a4996"),
				uuid.MustParse("c48ab7aa-08b9-405c-baf0-67e2dd382553"),
				uuid.MustParse("7142d1a0-28eb-480c-a23f-b9de2feb1fae"),
				uuid.MustParse("2ef8b796-d807-4c9c-b829-158f7fd459f7"),
				uuid.MustParse("bc46ae46-4562-488c-a0b9-041bafabd315"),
				uuid.MustParse("2227a9c3-503e-4e00-b816-d04d0f38404a"),
				uuid.MustParse("dacf034f-3c44-4bad-9494-3fcf0a5490bd"),
				uuid.MustParse("82e1b643-38c8-4a7f-bfba-3f94923a4f4a"),
				uuid.MustParse("234861a3-8402-415c-9474-94837dd00831"),
			},
			teachers: []uuid.UUID{
				uuid.MustParse("e4e880b0-f815-454d-8ff0-25cfe64905d0"),
				uuid.MustParse("2a264364-ab7c-4b21-8efe-9f10066dd9da"),
				uuid.MustParse("270a469c-af69-4bd1-a502-e34012a667e5"),
				uuid.MustParse("60c3edac-8327-4f67-bf99-1fac0994e03a"),
				uuid.MustParse("02ce4e77-bc76-4012-9389-5b1f0d3ced9f"),
				uuid.MustParse("78c6bf29-40b0-4aef-a1d8-c2c4c35f1731"),
				uuid.MustParse("d2ef64fb-b9e0-4c6b-8428-54ed4879529b"),
				uuid.MustParse("4a0f63ef-9f9f-4489-9feb-3a8d959687e0"),
				uuid.MustParse("16365040-8447-4119-bbd4-007b36048be7"),
				uuid.MustParse("b2425a49-2db5-4ab2-8c9d-6c7aeedb6c25"),
				uuid.MustParse("fca2350b-3bf0-49ef-8048-66e7086e8c02"),
				uuid.MustParse("688fe2b8-803e-44ac-977d-6325a9d11557"),
				uuid.MustParse("1f03a9d9-3fe8-4c47-be0c-33bd16d25f10"),
				uuid.MustParse("9c5789e0-bf19-4c00-9817-49d130e20704"),
				uuid.MustParse("738745df-0424-492c-ad07-2d8e30ca0197"),
				uuid.MustParse("ab335e2e-db68-444b-b5c9-db192f356159"),
				uuid.MustParse("7be6553d-f1f0-435d-860c-573238f620c0"),
				uuid.MustParse("9f3203eb-41fb-4502-b295-6ba143508655"),
				uuid.MustParse("563d1912-06f2-4351-abd8-0915fca2e3a4"),
				uuid.MustParse("ee5fc4e0-de3a-4fae-9363-054f9d762e21"),
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
