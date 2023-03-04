CREATE TABLE "users" (
                         "id" BIGSERIAL PRIMARY KEY,
                         "username" string UNIQUE NOT NULL,
                         "email" string UNIQUE NOT NULL,
                         "password" string NOT NULL
);

CREATE TABLE "papers" (
                          "arxiv_id" string PRIMARY KEY,
                          "title" string NOT NULL,
                          "abstract" string NOT NULL,
                          "authors" string NOT NULL,
                          "short_author" string NOT NULL,
                          "date" date NOT NULL
);

CREATE TABLE "saved_papers" (
                                "id" BIGSERIAL PRIMARY KEY,
                                "user_id" bigserial,
                                "paper_id" string
);

ALTER TABLE "saved_papers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "saved_papers" ADD FOREIGN KEY ("paper_id") REFERENCES "papers" ("arxiv_id");
