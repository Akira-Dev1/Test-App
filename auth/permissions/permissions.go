package permissions

// users
const (
	UserListRead      = "user:list:read"
	UserFullNameWrite = "user:fullName:write"
	UserDataRead      = "user:data:read"
	UserRolesRead     = "user:roles:read"
	UserRolesWrite    = "user:roles:write"
	UserBlockRead     = "user:block:read"
	UserBlockWrite    = "user:block:write"
)

// courses
const (
	CourseInfoWrite = "course:info:write"
	CourseTestList  = "course:testList"
	CourseTestRead  = "course:test:read"
	CourseTestWrite = "course:test:write"
	CourseTestAdd   = "course:test:add"
	CourseTestDel   = "course:test:del"
	CourseUserList  = "course:userList"
	CourseUserAdd   = "course:user:add"
	CourseUserDel   = "course:user:del"
	CourseAdd       = "course:add"
	CourseDel       = "course:del"
)

// questions
const (
	QuestListRead = "quest:list:read"
	QuestRead     = "quest:read"
	QuestUpdate   = "quest:update"
	QuestCreate   = "quest:create"
	QuestDel      = "quest:del"
)

// tests / answers
const (
	TestQuestAdd    = "test:quest:add"
	TestQuestDel    = "test:quest:del"
	TestQuestUpdate = "test:quest:update"
	TestAnswerRead  = "test:answer:read"
	AnswerRead      = "answer:read"
	AnswerUpdate    = "answer:update"
	AnswerDel       = "answer:del"
)
