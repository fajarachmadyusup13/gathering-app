-- gathering_app.attendees definition

CREATE TABLE `attendees` (
  `member_id` bigint NOT NULL,
  `gathering_id` bigint NOT NULL,
  `created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`deleted_at` DATETIME NULL,
  CONSTRAINT attendees_members_FK FOREIGN KEY (member_id) REFERENCES members(id),
  CONSTRAINT attendees_gatherings_FK FOREIGN KEY (gathering_id) REFERENCES gatherings(id)
);