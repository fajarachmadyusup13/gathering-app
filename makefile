SHELL:=/bin/bash

internal/model/mock/mock_member_repository.go:
	mockgen -destination=internal/model/mock/mock_member_repository.go -package=mock github.com/fajarachmadyusup13/gathering-app/internal/model MemberRepository
internal/model/mock/mock_invitation_repository.go:
	mockgen -destination=internal/model/mock/mock_invitation_repository.go -package=mock github.com/fajarachmadyusup13/gathering-app/internal/model InvitationRepository
internal/model/mock/mock_gathering_repository.go:
	mockgen -destination=internal/model/mock/mock_gathering_repository.go -package=mock github.com/fajarachmadyusup13/gathering-app/internal/model GatheringRepository
internal/model/mock/mock_attendee_repository.go:
	mockgen -destination=internal/model/mock/mock_attendee_repository.go -package=mock github.com/fajarachmadyusup13/gathering-app/internal/model AttendeeRepository
internal/model/mock/mock_member_usecase.go:
	mockgen -destination=internal/model/mock/mock_member_usecase.go -package=mock github.com/fajarachmadyusup13/gathering-app/internal/model MemberUsecase

mockgen: internal/model/mock/mock_member_repository.go \
	internal/model/mock/mock_invitation_repository.go \
	internal/model/mock/mock_gathering_repository.go \
	internal/model/mock/mock_attendee_repository.go \
	internal/model/mock/mock_member_usecase.go

clean:
	rm -v internal/model/mock/mock_*.go

check-modd-exists:
	@modd --version > /dev/null

run: check-modd-exists
	@modd -f ./.modd/server.modd.conf
	
