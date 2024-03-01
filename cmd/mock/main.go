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
				StudyPlaceId: uuid.MustParse("6d734a10-9bca-46c6-b9be-1d9a63170028"),
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
			from:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			to:         time.Date(2024, 2, 25, 0, 0, 0, 0, time.Local),
			groups: []uuid.UUID{
				uuid.MustParse("0ef94e37-273e-4a94-954e-447d490e7e3c"),
				uuid.MustParse("840fb256-0685-4d21-a1be-12ffb18bc0d3"),
				uuid.MustParse("882ca4b2-bc61-4659-b448-bdb063f68e33"),
				uuid.MustParse("e705d1b3-e947-4be3-aecd-618dcebd6a80"),
				uuid.MustParse("50d1f38c-dd99-438f-a44d-0f3ec1a688ec"),
				uuid.MustParse("eb8a35e2-76f6-48ff-8667-9db2d7513df1"),
				uuid.MustParse("62c2e1f2-d432-4f5f-aca4-efcbe2401849"),
				uuid.MustParse("02c69491-2184-427a-bbe6-08e1136b5c12"),
				uuid.MustParse("65f03052-fd95-49f5-8542-57e50e54e5e6"),
				uuid.MustParse("53ccf247-d7d4-4525-b79d-09c5b7a13445"),
			},
			rooms: []uuid.UUID{
				uuid.MustParse("5e821032-f799-42d5-9b90-8226602949da"),
				uuid.MustParse("ec0f6f58-40e5-4a20-8532-e7bc11bd285b"),
				uuid.MustParse("e176bc04-31e8-48a1-8a61-0471f6951560"),
				uuid.MustParse("766cef3e-fc55-4547-9548-f638fa7cfff9"),
				uuid.MustParse("e61868d9-7524-4df1-bcf8-9868dc2f8f49"),
				uuid.MustParse("c136cd6e-8462-4f19-a300-5c3d8be9a801"),
				uuid.MustParse("0607c8c4-8cfe-4925-bf6c-4fb5c1d6360c"),
				uuid.MustParse("4ad1e6c7-fce1-43f0-8235-3c45fba9dd0d"),
				uuid.MustParse("355b7a72-67df-4edd-a457-3f97d9809baf"),
				uuid.MustParse("4a9a5e87-009c-4d6f-bc6f-a0493854f753"),
				uuid.MustParse("fe6e975d-30db-40c1-967a-44dfc3486114"),
				uuid.MustParse("ec3446b6-f945-4f42-a657-5bafe5124440"),
				uuid.MustParse("5d5d775c-8b70-4e81-b33f-e324f2d96421"),
				uuid.MustParse("0942950a-8277-4bd2-a9f1-3e3e455ba69b"),
				uuid.MustParse("100f195f-abc9-467d-aac3-5e59339f4bc8"),
				uuid.MustParse("2f417772-5111-4775-8bf9-490510ed1387"),
				uuid.MustParse("70be8f06-7624-4260-97df-2fb96ef72268"),
				uuid.MustParse("b5984e5c-512f-4181-bfb7-dd698f11acb5"),
				uuid.MustParse("a3f3c2bf-8aac-45cf-9adf-0d6224636db3"),
				uuid.MustParse("cff72d91-c6db-45f1-ae55-69d3a12c77e2"),
				uuid.MustParse("63550dff-beb9-4200-a623-4ebcf8649dbf"),
				uuid.MustParse("7eb37a40-b7eb-45f7-aef0-43733b79e8d8"),
				uuid.MustParse("66e07f63-b892-44da-ad4d-12f669bb2ea7"),
				uuid.MustParse("9ea8db7a-5241-43b0-b4e1-2072737d8ece"),
				uuid.MustParse("ea05be3b-3891-4a0f-93bc-483d4627aaab"),
				uuid.MustParse("7df24c3a-91bf-46ec-a396-fad99794147f"),
				uuid.MustParse("8769405e-6430-43c2-8f70-b97683c1ef2a"),
				uuid.MustParse("f98229df-8448-4b57-af42-b27070f4dc6a"),
				uuid.MustParse("e89ea14d-5ad0-4751-a97d-60093dffbd64"),
				uuid.MustParse("c40fe6f6-4d32-4f9d-b04d-20e89a1ffa1e"),
				uuid.MustParse("a2607c2a-d7c9-4387-b2aa-528c252f9903"),
				uuid.MustParse("181d967d-81ce-462a-9767-32f415c70cd4"),
				uuid.MustParse("db2b52cd-cab4-485e-9ab9-5527bb225457"),
				uuid.MustParse("2977c49f-b07a-4ac7-a3ca-118a681efd8f"),
				uuid.MustParse("f6ca7922-9e68-448b-845e-1efdaad56c90"),
				uuid.MustParse("a4bdbc86-9471-42eb-a5ad-dc05e1bbde73"),
				uuid.MustParse("f8317ae6-2bad-4f48-9fb4-2afbd966a4c3"),
				uuid.MustParse("b11f2c52-2515-446e-8b41-500a9998699a"),
				uuid.MustParse("2e695245-4126-4272-8957-5c10248a7c6c"),
				uuid.MustParse("66b26fc6-dc5e-4ecd-ad61-67641ce90da9"),
				uuid.MustParse("51a5be92-490f-45cd-882b-e68440694d05"),
				uuid.MustParse("1454305d-068d-43a4-98a1-185ceac031dd"),
				uuid.MustParse("677a58c2-6233-457a-bb3d-9669e04aba95"),
				uuid.MustParse("e05b885f-9d2f-4968-a107-35b34b6d78b5"),
				uuid.MustParse("2da327e4-c213-4a9d-8a52-209864f9a758"),
				uuid.MustParse("a6af273c-8fed-4733-9495-133e00a0ee45"),
				uuid.MustParse("05fcbb97-ce05-4d24-b4ad-fc2e4c1ff8ac"),
				uuid.MustParse("7a5a6a32-451c-45ca-8966-8d2e28ee1be9"),
				uuid.MustParse("b03ad66f-75e7-4f23-9f02-56bac40b24ac"),
				uuid.MustParse("65d59e9a-dd94-4ce1-85e8-22385076711d"),
			},
			subjects: []uuid.UUID{
				uuid.MustParse("7a14dbbb-cab0-41b9-8806-992719aa9713"),
				uuid.MustParse("f2b83eb1-cf02-4517-b401-d88aa87b00e5"),
				uuid.MustParse("c870e076-d257-4d5a-9d1b-be1795ef6d2f"),
				uuid.MustParse("2965bf7b-ef4c-4def-9a0d-a92c1992c0d0"),
				uuid.MustParse("a322da57-9dd1-4130-9d74-e6a193f995d6"),
				uuid.MustParse("86dab03c-577a-43bd-808e-e1b14323de35"),
				uuid.MustParse("0fe3ed16-792f-450f-98fe-197434d442cf"),
				uuid.MustParse("e50304c3-681e-45bd-9b13-a06d242bbf40"),
				uuid.MustParse("5af34c76-3b6b-416a-91c8-c7cfb4dee53c"),
				uuid.MustParse("85bcc1a4-ff87-4be5-a298-398b73a51b7d"),
				uuid.MustParse("16e4fd95-0f32-4845-aaf3-61b53fabe5ad"),
				uuid.MustParse("0682205f-29fb-49ee-84c8-1304c6ef2bbd"),
				uuid.MustParse("49b5eff8-0f07-4c93-8bff-c8d51a5dc7e1"),
				uuid.MustParse("e18b7370-a998-466a-8c9f-605390a6b964"),
				uuid.MustParse("94aa0dc8-71e0-44a2-be24-088869e4ce63"),
				uuid.MustParse("cc0233d7-efe5-4876-9e16-e1acd43a1bd0"),
				uuid.MustParse("8a2a78a3-3d27-4f99-8437-277ebb36630a"),
				uuid.MustParse("c40341ab-ced1-4142-98b2-54915a817950"),
				uuid.MustParse("3f414ee0-9e8e-4f32-af79-8e404db5fb3d"),
				uuid.MustParse("62212092-6b23-4b11-8aab-8303f554170f"),
				uuid.MustParse("94bab289-22c5-41a9-8ffa-de43e89c6159"),
				uuid.MustParse("2d8aaa20-ff57-492b-8c33-61b7f4b85973"),
				uuid.MustParse("17ef6543-fcae-40a7-a248-79844e0bf4d3"),
				uuid.MustParse("1d64633e-6447-48bd-b852-fd2b14c1ec93"),
				uuid.MustParse("0d29dac1-ee02-4d32-9888-50896ac3153c"),
				uuid.MustParse("07e64cdf-25fe-4935-abae-01d9cb9bca9a"),
				uuid.MustParse("90bd525b-04bd-43f7-87dd-8a24ece74b4e"),
				uuid.MustParse("8714c156-7c20-4ec8-ae33-c6955a39b971"),
				uuid.MustParse("c4c36582-17c5-4d04-acf8-255394070c63"),
				uuid.MustParse("16061287-193a-430a-8ebd-fb70c16239ca"),
				uuid.MustParse("8ff5d0ea-a7d2-4afc-a6ff-1a36f3bbf08b"),
				uuid.MustParse("f65d5b58-d588-4270-9111-80e8ad80253a"),
				uuid.MustParse("beb53eb7-bece-4a2f-8461-45f2a83bbc5e"),
				uuid.MustParse("5794df7e-c07e-4f75-b565-cea3473c771c"),
				uuid.MustParse("be027066-1553-45ad-8fc2-ab96ca33614f"),
				uuid.MustParse("8a1a6efc-d6bc-4a66-8ec1-0c7c746feae6"),
				uuid.MustParse("01f02c43-9226-48c0-ae4b-cb73e9226e94"),
				uuid.MustParse("4223f959-d0ea-4742-9071-9a31df3970a5"),
				uuid.MustParse("e7f73269-c0ec-456b-a51f-b15489bfedbe"),
				uuid.MustParse("6eca473f-f4b3-4e73-b85f-9e59062fca59"),
			},
			teachers: []uuid.UUID{
				uuid.MustParse("f28dbca4-7abf-4847-85a1-0232932646c6"),
				uuid.MustParse("25f31e15-3535-4651-922c-967dc06519c8"),
				uuid.MustParse("d6330622-9761-4bc6-8d47-3be80e8b24b6"),
				uuid.MustParse("060a70dd-577a-43a3-b615-abe28965fe7b"),
				uuid.MustParse("168c06d8-1a2f-493d-8b3d-a45a320b2e0d"),
				uuid.MustParse("b9ca5dcc-0b9d-4b18-ab64-479ffaea771d"),
				uuid.MustParse("a6678cf5-3f4c-44fb-a45d-ebe275fb175c"),
				uuid.MustParse("428f3721-5c7b-4e0c-9d39-cc556617517a"),
				uuid.MustParse("98cf64da-fb4e-4fc9-97ef-6fe219623cae"),
				uuid.MustParse("879c69f1-2cda-4cc8-b221-52011920f5aa"),
				uuid.MustParse("020d6efb-f2db-49cd-ad6d-447a90f96c0d"),
				uuid.MustParse("d92ed7ef-4c46-49f1-b518-dbf6c3611b15"),
				uuid.MustParse("f7957166-ac07-4797-856d-7eeb5ff3d1f5"),
				uuid.MustParse("9f84ff12-6e5f-44da-96f0-c6fc535a9388"),
				uuid.MustParse("b435af49-d4d3-4cbe-84a9-7f9a7ea87e56"),
				uuid.MustParse("18400364-997b-47b3-bba4-9e0b4d68ca7f"),
				uuid.MustParse("5da5213f-73ea-4762-8703-d0ec77b6d72d"),
				uuid.MustParse("cfbdf9de-84a4-4ae1-b9ec-9ad425d3acf7"),
				uuid.MustParse("2df7cc10-5cdd-45af-9a0a-b1f5a69be840"),
				uuid.MustParse("b17dae31-1757-4a9b-8a03-bac6ca874d1e"),
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
