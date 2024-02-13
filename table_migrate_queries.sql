CREATE TABLE "teachers" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "students" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "class_groups" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "classes" (
  "id" bigserial PRIMARY KEY,
  "class_groups_id" varchar,
  "name" varchar NOT NULL
);

CREATE TABLE "student_class" (
  "id" bigserial PRIMARY KEY,
  "student_id" varchar,
  "class_id" varchar
);

CREATE TABLE "subjects" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "subject_reports" (
  "id" bigserial PRIMARY KEY,
  "subject_id" bigint,
  "teacher_id" bigint,
  "student_class_id" bigint,
  "mid_exam_result" int,
  "final_exam_result" int,
  "seasonal_exam_result" int[]
);

COMMENT ON COLUMN "teachers"."id" IS 'from nip';

COMMENT ON COLUMN "students"."id" IS 'from nis';

COMMENT ON COLUMN "class_groups"."id" IS 'kode jurusan';

ALTER TABLE "classes" ADD FOREIGN KEY ("class_groups_id") REFERENCES "class_groups" ("id");

ALTER TABLE "student_class" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "student_class" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "subject_reports" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "subject_reports" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");

ALTER TABLE "subject_reports" ADD FOREIGN KEY ("student_class_id") REFERENCES "student_class" ("id");
