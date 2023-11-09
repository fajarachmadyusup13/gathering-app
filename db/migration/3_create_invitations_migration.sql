CREATE TABLE gathering_app.invitations (
	id BIGINT NOT NULL,
	member_id BIGINT NOT NULL,
	gathering_id BIGINT NOT NULL,
	status INT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	deleted_at DATETIME NULL,
    PRIMARY KEY (id),
	CONSTRAINT invitations_members_FK FOREIGN KEY (member_id) REFERENCES members(id),
    CONSTRAINT invitations_gatherings_FK FOREIGN KEY (gathering_id) REFERENCES gatherings(id)
);