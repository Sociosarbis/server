package model

const (
	RevisionTypeSubject                  = 1  // 条目
	RevisionTypeSubjectCharacterRelation = 5  // 条目->角色关联
	RevisionTypeSubjectCastRelation      = 6  // 条目->声优关联
	RevisionTypeSubjectPersonRelation    = 10 // 条目->人物关联
	RevisionTypeSubjectMerge             = 11 // 条目管理
	RevisionTypeSubjectErase             = 12
	RevisionTypeSubjectRelation          = 17 // 条目关联
	RevisionTypeSubjectLock              = 103
	RevisionTypeSubjectUnlock            = 104
	RevisionTypeCharacter                = 2  // 角色
	RevisionTypeCharacterSubjectRelation = 4  // 角色->条目关联
	RevisionTypeCharacterCastRelation    = 7  // 角色->声优关联
	RevisionTypeCharacterMerge           = 13 // 角色管理
	RevisionTypeCharacterErase           = 14
	RevisionTypePerson                   = 3  // 人物
	RevisionTypePersonCastRelation       = 8  // 人物->声优关联
	RevisionTypePersonSubjectRelation    = 9  // 人物->条目关联
	RevisionTypePersonMerge              = 15 // 人物管理
	RevisionTypePersonErase              = 16
	RevisionTypeEp                       = 18  // 章节
	RevisionTypeEpMerge                  = 181 // 章节管理
	RevisionTypeEpMove                   = 182
	RevisionTypeEpLock                   = 183
	RevisionTypeEpUnlock                 = 184
	RevisionTypeEpErase                  = 185
)

var PersonRevisionTypes = []uint8{
	RevisionTypePerson,
	RevisionTypePersonCastRelation,
	RevisionTypePersonErase,
	RevisionTypePersonMerge,
	RevisionTypePersonSubjectRelation,
}
