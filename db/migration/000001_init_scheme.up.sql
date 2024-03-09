CREATE TABLE "teachers" (
  "id" bigserial PRIMARY KEY,
  "nip" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE "students" (
  "id" bigserial PRIMARY KEY,
  "nis" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE "class_groups" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "classes" (
  "id" bigserial PRIMARY KEY,
  "class_groups_id" varchar NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "student_classes" (
  "id" bigserial PRIMARY KEY,
  "student_id" bigint NOT NULL,
  "class_id" bigint NOT NULL,
  "period" varchar NOT NULL
);

CREATE TABLE "subjects" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "teacher_subjects" (
  "id" bigserial PRIMARY KEY,
  "teacher_id" bigint NOT NULL,
  "subject_id" bigint NOT NULL,
  "period" varchar NOT NULL
);

CREATE TABLE "subject_reports" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz DEFAULT (now()) NOT NULL,
  "teacher_subject_id" bigint NOT NULL,
  "student_class_id" bigint NOT NULL,
  "mid_exam_result" int,
  "final_exam_result" int,
  "monthly_exam_result" int[]
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz DEFAULT (now()) NOT NULL,
  "name" varchar NOT NULL,
  "role" varchar NOT NULL
);

COMMENT ON COLUMN "class_groups"."id" IS 'kode jurusan';

COMMENT ON COLUMN "student_classes"."period" IS 'tahun awal TA';

COMMENT ON COLUMN "teacher_subjects"."period" IS 'tahun awal TA';

COMMENT ON COLUMN "users"."role" IS 'enum teacher, student';

ALTER TABLE "teachers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "students" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "classes" ADD FOREIGN KEY ("class_groups_id") REFERENCES "class_groups" ("id");

ALTER TABLE "student_classes" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "student_classes" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "teacher_subjects" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");

ALTER TABLE "teacher_subjects" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "subject_reports" ADD FOREIGN KEY ("teacher_subject_id") REFERENCES "teacher_subjects" ("id");

ALTER TABLE "subject_reports" ADD FOREIGN KEY ("student_class_id") REFERENCES "student_classes" ("id");
