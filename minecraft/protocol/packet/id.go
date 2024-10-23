package packet

const (
	IDLogin = iota + 1
	IDPlayStatus
	IDServerToClientHandshake
	IDClientToServerHandshake
	IDDisconnect
	IDResourcePacksInfo
	IDResourcePackStack
	IDResourcePackClientResponse
	IDText
	IDSetTime
	IDStartGame
	IDAddPlayer
	IDAddActor
	IDRemoveActor
	IDAddItemActor
	_
	IDTakeItemActor
	IDMoveActorAbsolute
	IDMovePlayer
	IDPassengerJump
	IDUpdateBlock
	IDAddPainting
	IDTickSync
	_
	IDLevelEvent
	IDBlockEvent
	IDActorEvent
	IDMobEffect
	IDUpdateAttributes
	IDInventoryTransaction
	IDMobEquipment
	IDMobArmourEquipment
	IDInteract
	IDBlockPickRequest
	IDActorPickRequest
	IDPlayerAction
	_
	IDHurtArmour
	IDSetActorData
	IDSetActorMotion
	IDSetActorLink
	IDSetHealth
	IDSetSpawnPosition
	IDAnimate
	IDRespawn
	IDContainerOpen
	IDContainerClose
	IDPlayerHotBar
	IDInventoryContent
	IDInventorySlot
	IDContainerSetData
	IDCraftingData
	_
	IDGUIDataPickItem
	IDAdventureSettings
	IDBlockActorData
	IDPlayerInput
	IDLevelChunk
	IDSetCommandsEnabled
	IDSetDifficulty
	IDChangeDimension
	IDSetPlayerGameType
	IDPlayerList
	IDSimpleEvent
	IDEvent
	IDSpawnExperienceOrb
	IDClientBoundMapItemData
	IDMapInfoRequest
	IDRequestChunkRadius
	IDChunkRadiusUpdated
	_
	IDGameRulesChanged
	IDCamera
	IDBossEvent
	IDShowCredits
	IDAvailableCommands
	IDCommandRequest
	IDCommandBlockUpdate
	IDCommandOutput
	IDUpdateTrade
	IDUpdateEquip
	IDResourcePackDataInfo
	IDResourcePackChunkData
	IDResourcePackChunkRequest
	IDTransfer
	IDPlaySound
	IDStopSound
	IDSetTitle
	IDAddBehaviourTree
	IDStructureBlockUpdate
	IDShowStoreOffer
	IDPurchaseReceipt
	IDPlayerSkin
	IDSubClientLogin
	IDAutomationClientConnect
	IDSetLastHurtBy
	IDBookEdit
	IDNPCRequest
	IDPhotoTransfer
	IDModalFormRequest
	IDModalFormResponse
	IDServerSettingsRequest
	IDServerSettingsResponse
	IDShowProfile
	IDSetDefaultGameType
	IDRemoveObjective
	IDSetDisplayObjective
	IDSetScore
	IDLabTable
	IDUpdateBlockSynced
	IDMoveActorDelta
	IDSetScoreboardIdentity
	IDSetLocalPlayerAsInitialised
	IDUpdateSoftEnum
	IDNetworkStackLatency
	_
	IDScriptCustomEvent
	IDSpawnParticleEffect
	IDAvailableActorIdentifiers
	_
	IDNetworkChunkPublisherUpdate
	IDBiomeDefinitionList
	IDLevelSoundEvent
	IDLevelEventGeneric
	IDLecternUpdate
	_
	_
	_
	IDClientCacheStatus
	IDOnScreenTextureAnimation
	IDMapCreateLockedCopy
	IDStructureTemplateDataRequest
	IDStructureTemplateDataResponse
	_
	IDClientCacheBlobStatus
	IDClientCacheMissResponse
	IDEducationSettings
	IDEmote
	IDMultiPlayerSettings
	IDSettingsCommand
	IDAnvilDamage
	IDCompletedUsingItem
	IDNetworkSettings
	IDPlayerAuthInput
	IDCreativeContent
	IDPlayerEnchantOptions
	IDItemStackRequest
	IDItemStackResponse
	IDPlayerArmourDamage
	IDCodeBuilder
	IDUpdatePlayerGameType
	IDEmoteList
	IDPositionTrackingDBServerBroadcast
	IDPositionTrackingDBClientRequest
	IDDebugInfo
	IDPacketViolationWarning
	IDMotionPredictionHints
	IDAnimateEntity
	IDCameraShake
	IDPlayerFog
	IDCorrectPlayerMovePrediction
	IDItemComponent
	IDFilterText
	IDClientBoundDebugRenderer
	IDSyncActorProperty
	IDAddVolumeEntity
	IDRemoveVolumeEntity
	IDSimulationType
	IDNPCDialogue
	IDEducationResourceURI
	IDCreatePhoto
	IDUpdateSubChunkBlocks
	IDPhotoInfoRequest
	IDSubChunk
	IDSubChunkRequest
	IDClientStartItemCooldown
	IDScriptMessage
	IDCodeBuilderSource
	IDTickingAreasLoadStatus
	IDDimensionData
	IDAgentAction
	IDChangeMobProperty
	IDLessonProgress
	IDRequestAbility
	IDRequestPermissions
	IDToastRequest
	IDUpdateAbilities
	IDUpdateAdventureSettings
	IDDeathInfo
	IDEditorNetwork
	IDFeatureRegistry
	IDServerStats
	IDRequestNetworkSettings
	IDGameTestRequest
	IDGameTestResults
	IDUpdateClientInputLocks
	IDClientCheatAbility
	IDCameraPresets
	IDUnlockedRecipes
	IDCameraInstruction = iota + 101
	IDCompressedBiomeDefinitionList
	IDTrimData
	IDOpenSign
	IDAgentAnimation
	IDRefreshEntitlements
	IDPlayerToggleCrafterSlotRequest
	IDSetPlayerInventoryOptions
	IDSetHud
	IDAwardAchievement
	IDClientBoundCloseForm
	_
	IDServerBoundLoadingScreen
	IDJigsawStructureData
	IDCurrentStructureFeature
	IDServerBoundDiagnostics
	IDCameraAimAssist
	IDContainerRegistryCleanup
	IDMovementEffect
	IDSetMovementAuthority
)
