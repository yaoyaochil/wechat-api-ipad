package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["wechatwebapi/controllers:FavorController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FavorController"],
        beego.ControllerComments{
            Method: "DelFavItem",
            Router: `/DelFavItem`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FavorController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FavorController"],
        beego.ControllerComments{
            Method: "FavSync",
            Router: `/FavSync`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FavorController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FavorController"],
        beego.ControllerComments{
            Method: "GetFavInfo",
            Router: `/GetFavInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FavorController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FavorController"],
        beego.ControllerComments{
            Method: "GetFavItem",
            Router: `/GetFavItem`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SetSnsViewTime",
            Router: `/SetSnsViewTime`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SnsComment",
            Router: `/SnsComment`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SnsObjectDetail",
            Router: `/SnsObjectDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SnsObjectTop",
            Router: `/SnsObjectTop`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SnsPost",
            Router: `/SnsPost`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SnsSync",
            Router: `/SnsSync`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SnsTimeline",
            Router: `/SnsTimeline`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SnsUpload",
            Router: `/SnsUpload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendCircleController"],
        beego.ControllerComments{
            Method: "SnsUserPage",
            Router: `/SnsUserPage`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "DelFriend",
            Router: `/DelFriend`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "DeleteFromContact",
            Router: `/DeleteFromContact`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "GetContact",
            Router: `/GetContact`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "GetContactList",
            Router: `/GetContactList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "GetContactStatus",
            Router: `/GetContactStatus`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "GetMFriend",
            Router: `/GetMFriend`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "SearchContact",
            Router: `/SearchContact`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "SetContactBlacklist",
            Router: `/SetContactBlacklist`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "SetContactRemarks",
            Router: `/SetContactRemarks`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "UploadMContact",
            Router: `/UploadMContact`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:FriendController"],
        beego.ControllerComments{
            Method: "VerifyUser",
            Router: `/VerifyUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "AddChatRoomMember",
            Router: `/AddChatRoomMember`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "CreateChatRoom",
            Router: `/CreateChatRoom`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "DelChatRoomMember",
            Router: `/DelChatRoomMember`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetChatRoomInfo",
            Router: `/GetChatRoomInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetChatRoomInfoDetail",
            Router: `/GetChatRoomInfoDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetChatRoomMemberDetail",
            Router: `/GetChatRoomMemberDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "InviteChatRoomMember",
            Router: `/InviteChatRoomMember`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "MoveToContract",
            Router: `/MoveToContract`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "OperateChatRoomAdmin",
            Router: `/OperateChatRoomAdmin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "QuitChatroom",
            Router: `/QuitChatroom`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "ScanIntoGroup",
            Router: `/ScanIntoGroup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SetChatRoomAnnouncement",
            Router: `/SetChatRoomAnnouncement`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SetChatRoomName",
            Router: `/SetChatRoomName`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SetChatRoomRemarks",
            Router: `/SetChatRoomRemarks`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"],
        beego.ControllerComments{
            Method: "AddContactLabel",
            Router: `/AddContactLabel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"],
        beego.ControllerComments{
            Method: "DeleteContactLabel",
            Router: `/DeleteContactLabel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"],
        beego.ControllerComments{
            Method: "GetContactLabelList",
            Router: `/GetContactLabelList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"],
        beego.ControllerComments{
            Method: "UpdateContactLabel",
            Router: `/UpdateContactLabel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LabelController"],
        beego.ControllerComments{
            Method: "UpdateContactLabelList",
            Router: `/UpdateContactLabelList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "AutoAuth",
            Router: `/AutoAuth`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Awaken",
            Router: `/Awaken`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "CheckLoginQrcode",
            Router: `/CheckLoginQrcode`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Get62Data",
            Router: `/Get62Data`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "GetCacheInfo",
            Router: `/GetCacheInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "GetLoginQrcode",
            Router: `/GetLoginQrcode`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "HeartBeat",
            Router: `/HeartBeat`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "LogOut",
            Router: `/LogOut`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Manualauth",
            Router: `/Manualauth`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Newinit",
            Router: `/Newinit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "MassSend",
            Router: `/MassSend`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "NewSync",
            Router: `/NewSync`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "RevokeMsg",
            Router: `/RevokeMsg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "SendAppMsg",
            Router: `/SendAppMsg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "SendImg",
            Router: `/SendImg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "SendMsg",
            Router: `/SendMsg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "SendVideo",
            Router: `/SendVideo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "SendVoice",
            Router: `/SendVoice`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "ShareCard",
            Router: `/ShareCard`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:MsgController"],
        beego.ControllerComments{
            Method: "ShareLocation",
            Router: `/ShareLocation`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"],
        beego.ControllerComments{
            Method: "BindMobile",
            Router: `/BindMobile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"],
        beego.ControllerComments{
            Method: "DelSafetyInfo",
            Router: `/DelSafetyInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"],
        beego.ControllerComments{
            Method: "DownloadVoice",
            Router: `/DownloadVoice`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"],
        beego.ControllerComments{
            Method: "GetA8Key",
            Router: `/GetA8Key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"],
        beego.ControllerComments{
            Method: "GetQrcode",
            Router: `/GetQrcode`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"],
        beego.ControllerComments{
            Method: "GetSafetyInfo",
            Router: `/GetSafetyInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"],
        beego.ControllerComments{
            Method: "MPGetA8Key",
            Router: `/MPGetA8Key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:ToolsController"],
        beego.ControllerComments{
            Method: "SetProxy",
            Router: `/setproxy`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetProfile",
            Router: `/GetProfile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "SetPassword",
            Router: `/SetPassword`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateProfile",
            Router: `/UpdateProfile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "UploadHeadImage",
            Router: `/UploadHeadImage`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "VerifyPassword",
            Router: `/VerifyPassword`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatwebapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "VerifySwitch",
            Router: `/VerifySwitch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
