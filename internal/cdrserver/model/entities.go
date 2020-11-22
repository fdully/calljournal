package model

import (
	"encoding/xml"
	"net/url"
	"strconv"

	"github.com/google/uuid"
)

type CallPath struct {
	UUID uuid.UUID `json:"uuid"`
	DIRC string    `json:"dirc"`
	YEAR string    `json:"year"`
	MONT string    `json:"mont"`
	DAYX string    `json:"dayx"`
	NAME string    `json:"name"`
}

type CDR struct {
	XMLName     xml.Name `xml:"cdr"`
	Text        string   `xml:",chardata"`
	CoreUUID    string   `xml:"core-uuid,attr"`
	Switchname  string   `xml:"switchname,attr"`
	ChannelData struct {
		Text        string `xml:",chardata"`
		State       string `xml:"state"`
		Direction   string `xml:"direction"`
		StateNumber string `xml:"state_number"`
		Flags       string `xml:"flags"`
		Caps        string `xml:"caps"`
	} `xml:"channel_data"`
	CallStats struct {
		Text  string `xml:",chardata"`
		Audio struct {
			Text    string `xml:",chardata"`
			Inbound struct {
				Text              string `xml:",chardata"`
				RawBytes          string `xml:"raw_bytes"`
				MediaBytes        string `xml:"media_bytes"`
				PacketCount       string `xml:"packet_count"`
				MediaPacketCount  string `xml:"media_packet_count"`
				SkipPacketCount   string `xml:"skip_packet_count"`
				JitterPacketCount string `xml:"jitter_packet_count"`
				DtmfPacketCount   string `xml:"dtmf_packet_count"`
				CngPacketCount    string `xml:"cng_packet_count"`
				FlushPacketCount  string `xml:"flush_packet_count"`
				LargestJbSize     string `xml:"largest_jb_size"`
				JitterMinVariance string `xml:"jitter_min_variance"`
				JitterMaxVariance string `xml:"jitter_max_variance"`
				JitterLossRate    string `xml:"jitter_loss_rate"`
				JitterBurstRate   string `xml:"jitter_burst_rate"`
				MeanInterval      string `xml:"mean_interval"`
				FlawTotal         string `xml:"flaw_total"`
				QualityPercentage string `xml:"quality_percentage"`
				Mos               string `xml:"mos"`
			} `xml:"inbound"`
			Outbound struct {
				Text             string `xml:",chardata"`
				RawBytes         string `xml:"raw_bytes"`
				MediaBytes       string `xml:"media_bytes"`
				PacketCount      string `xml:"packet_count"`
				MediaPacketCount string `xml:"media_packet_count"`
				SkipPacketCount  string `xml:"skip_packet_count"`
				DtmfPacketCount  string `xml:"dtmf_packet_count"`
				CngPacketCount   string `xml:"cng_packet_count"`
				RtcpPacketCount  string `xml:"rtcp_packet_count"`
				RtcpOctetCount   string `xml:"rtcp_octet_count"`
			} `xml:"outbound"`
			ErrorLog struct {
				Text        string `xml:",chardata"`
				ErrorPeriod struct {
					Text             string `xml:",chardata"`
					Start            string `xml:"start"`
					Stop             string `xml:"stop"`
					Flaws            string `xml:"flaws"`
					ConsecutiveFlaws string `xml:"consecutive-flaws"`
					DurationMsec     string `xml:"duration-msec"`
				} `xml:"error-period"`
			} `xml:"error-log"`
		} `xml:"audio"`
	} `xml:"call-stats"`
	Variables struct {
		Text                        string    `xml:",chardata"`
		Direction                   string    `xml:"direction"`
		UUID                        string    `xml:"uuid"`
		SessionID                   string    `xml:"session_id"`
		SIPFromParams               string    `xml:"sip_from_params"`
		SIPFromUser                 string    `xml:"sip_from_user"`
		SIPFromURI                  string    `xml:"sip_from_uri"`
		SIPFromHost                 string    `xml:"sip_from_host"`
		VideoMediaFlow              string    `xml:"video_media_flow"`
		TextMediaFlow               string    `xml:"text_media_flow"`
		ChannelName                 string    `xml:"channel_name"`
		SIPLocalNetworkAddr         string    `xml:"sip_local_network_addr"`
		SIPNetworkIP                string    `xml:"sip_network_ip"`
		SIPNetworkPort              string    `xml:"sip_network_port"`
		SIPInviteStamp              string    `xml:"sip_invite_stamp"`
		SIPReceivedIP               string    `xml:"sip_received_ip"`
		SIPReceivedPort             string    `xml:"sip_received_port"`
		SIPViaProtocol              string    `xml:"sip_via_protocol"`
		SIPAuthorized               string    `xml:"sip_authorized"`
		SIPAclAuthedBy              string    `xml:"sip_acl_authed_by"`
		SIPFromUserStripped         string    `xml:"sip_from_user_stripped"`
		SIPFromEpid                 string    `xml:"sip_from_epid"`
		SofiaProfileName            string    `xml:"sofia_profile_name"`
		SofiaProfileURL             string    `xml:"sofia_profile_url"`
		RecoveryProfileName         string    `xml:"recovery_profile_name"`
		SIPPAssertedIdentity        string    `xml:"sip_P-Asserted-Identity"`
		SIPCidType                  string    `xml:"sip_cid_type"`
		SIPAllow                    string    `xml:"sip_allow"`
		SIPReqParams                string    `xml:"sip_req_params"`
		SIPReqUser                  string    `xml:"sip_req_user"`
		SIPReqURI                   string    `xml:"sip_req_uri"`
		SIPReqHost                  string    `xml:"sip_req_host"`
		SIPToParams                 string    `xml:"sip_to_params"`
		SIPToUser                   string    `xml:"sip_to_user"`
		SIPToURI                    string    `xml:"sip_to_uri"`
		SIPToHost                   string    `xml:"sip_to_host"`
		SIPContactParams            string    `xml:"sip_contact_params"`
		SIPContactUser              string    `xml:"sip_contact_user"`
		SIPContactPort              string    `xml:"sip_contact_port"`
		SIPContactURI               string    `xml:"sip_contact_uri"`
		SIPContactHost              string    `xml:"sip_contact_host"`
		SIPViaHost                  string    `xml:"sip_via_host"`
		SIPViaPort                  string    `xml:"sip_via_port"`
		SIPViaRport                 string    `xml:"sip_via_rport"`
		PresenceID                  string    `xml:"presence_id"`
		SIPPrivacy                  string    `xml:"sip_Privacy"`
		SwitchRSdp                  string    `xml:"switch_r_sdp"`
		EpCodecString               string    `xml:"ep_codec_string"`
		TransferHistory             []string  `xml:"transfer_history"`
		MediaBugAnswerReq           string    `xml:"media_bug_answer_req"`
		RECORDSTEREO                string    `xml:"RECORD_STEREO"`
		RTPUseCodecString           string    `xml:"rtp_use_codec_string"`
		RemoteAudioMediaFlow        string    `xml:"remote_audio_media_flow"`
		AudioMediaFlow              string    `xml:"audio_media_flow"`
		RTPRemoteAudioRtcpPort      string    `xml:"rtp_remote_audio_rtcp_port"`
		RTPAudioRecvPt              string    `xml:"rtp_audio_recv_pt"`
		RTPUseCodecName             string    `xml:"rtp_use_codec_name"`
		RTPUseCodecRate             string    `xml:"rtp_use_codec_rate"`
		RTPUseCodecPtime            string    `xml:"rtp_use_codec_ptime"`
		RTPUseCodecChannels         string    `xml:"rtp_use_codec_channels"`
		RTPLastAudioCodecString     string    `xml:"rtp_last_audio_codec_string"`
		ReadCodec                   string    `xml:"read_codec"`
		OriginalReadCodec           string    `xml:"original_read_codec"`
		ReadRate                    string    `xml:"read_rate"`
		OriginalReadRate            string    `xml:"original_read_rate"`
		WriteCodec                  string    `xml:"write_codec"`
		WriteRate                   string    `xml:"write_rate"`
		DtmfType                    string    `xml:"dtmf_type"`
		LocalMediaIP                string    `xml:"local_media_ip"`
		LocalMediaPort              string    `xml:"local_media_port"`
		AdvertisedMediaIP           string    `xml:"advertised_media_ip"`
		RTPUseTimerName             string    `xml:"rtp_use_timer_name"`
		RTPUsePt                    string    `xml:"rtp_use_pt"`
		RTPUseSsrc                  string    `xml:"rtp_use_ssrc"`
		RTP2833SendPayload          string    `xml:"rtp_2833_send_payload"`
		RTP2833RecvPayload          string    `xml:"rtp_2833_recv_payload"`
		RemoteMediaIP               string    `xml:"remote_media_ip"`
		RemoteMediaPort             string    `xml:"remote_media_port"`
		SessionInHangupHook         string    `xml:"session_in_hangup_hook"`
		APIHangupHook               string    `xml:"api_hangup_hook"`
		OriginationPrivacy          string    `xml:"origination_privacy"`
		Privacy                     string    `xml:"privacy"`
		ExportVars                  string    `xml:"export_vars"`
		X                           string    `xml:"x"`
		MaxForwards                 string    `xml:"max_forwards"`
		TransferSource              string    `xml:"transfer_source"`
		DPMATCH                     []string  `xml:"DP_MATCH"`
		CallUUID                    string    `xml:"call_uuid"`
		EffectiveCallerIDName       string    `xml:"effective_caller_id_name"`
		EffectiveCallerIDNumber     string    `xml:"effective_caller_id_number"`
		HangupAfterBridge           string    `xml:"hangup_after_bridge"`
		CallTimeout                 string    `xml:"call_timeout"`
		CurrentApplicationData      string    `xml:"current_application_data"`
		CurrentApplication          string    `xml:"current_application"`
		OriginatedLegs              string    `xml:"originated_legs"`
		OriginateDisposition        string    `xml:"originate_disposition"`
		DIALSTATUS                  string    `xml:"DIALSTATUS"`
		OriginateCauses             string    `xml:"originate_causes"`
		LastBridgeTo                string    `xml:"last_bridge_to"`
		BridgeChannel               string    `xml:"bridge_channel"`
		BridgeUUID                  string    `xml:"bridge_uuid"`
		SignalBond                  string    `xml:"signal_bond"`
		SwitchMSdp                  string    `xml:"switch_m_sdp"`
		RTPLocalSdpStr              string    `xml:"rtp_local_sdp_str"`
		EndpointDisposition         string    `xml:"endpoint_disposition"`
		SIPToTag                    string    `xml:"sip_to_tag"`
		SIPFromTag                  string    `xml:"sip_from_tag"`
		SIPCseq                     string    `xml:"sip_cseq"`
		SIPCallID                   string    `xml:"sip_call_id"`
		SIPFullVia                  string    `xml:"sip_full_via"`
		SIPRecoverVia               string    `xml:"sip_recover_via"`
		SIPFullFrom                 string    `xml:"sip_full_from"`
		SIPFullTo                   string    `xml:"sip_full_to"`
		LastSentCalleeIDName        string    `xml:"last_sent_callee_id_name"`
		LastSentCalleeIDNumber      string    `xml:"last_sent_callee_id_number"`
		SIPTermStatus               string    `xml:"sip_term_status"`
		ProtoSpecificHangupCause    string    `xml:"proto_specific_hangup_cause"`
		SIPTermCause                string    `xml:"sip_term_cause"`
		LastBridgeRole              string    `xml:"last_bridge_role"`
		SIPUserAgent                string    `xml:"sip_user_agent"`
		SIPHangupDisposition        string    `xml:"sip_hangup_disposition"`
		BridgeHangupCause           string    `xml:"bridge_hangup_cause"`
		RecordFileSize              xmlInt64  `xml:"record_file_size"`
		RecordSamples               string    `xml:"record_samples"`
		RecordSeconds               xmlInt64  `xml:"record_seconds"`
		RecordMs                    string    `xml:"record_ms"`
		RecordCompletionCause       string    `xml:"record_completion_cause"`
		RECORDDISCARDED             string    `xml:"RECORD_DISCARDED"`
		HangupCause                 string    `xml:"hangup_cause"`
		HangupCauseQ850             string    `xml:"hangup_cause_q850"`
		DigitsDialed                string    `xml:"digits_dialed"`
		StartStamp                  xmlString `xml:"start_stamp"`
		ProfileStartStamp           string    `xml:"profile_start_stamp"`
		AnswerStamp                 string    `xml:"answer_stamp"`
		BridgeStamp                 string    `xml:"bridge_stamp"`
		ProgressStamp               string    `xml:"progress_stamp"`
		ProgressMediaStamp          string    `xml:"progress_media_stamp"`
		EndStamp                    string    `xml:"end_stamp"`
		StartEpoch                  xmlInt64  `xml:"start_epoch"`
		StartUepoch                 string    `xml:"start_uepoch"`
		ProfileStartEpoch           string    `xml:"profile_start_epoch"`
		ProfileStartUepoch          string    `xml:"profile_start_uepoch"`
		AnswerEpoch                 xmlInt64  `xml:"answer_epoch"`
		AnswerUepoch                string    `xml:"answer_uepoch"`
		BridgeEpoch                 string    `xml:"bridge_epoch"`
		BridgeUepoch                string    `xml:"bridge_uepoch"`
		LastHoldEpoch               string    `xml:"last_hold_epoch"`
		LastHoldUepoch              string    `xml:"last_hold_uepoch"`
		HoldAccumSeconds            string    `xml:"hold_accum_seconds"`
		HoldAccumUsec               string    `xml:"hold_accum_usec"`
		HoldAccumMs                 string    `xml:"hold_accum_ms"`
		ResurrectEpoch              string    `xml:"resurrect_epoch"`
		ResurrectUepoch             string    `xml:"resurrect_uepoch"`
		ProgressEpoch               string    `xml:"progress_epoch"`
		ProgressUepoch              string    `xml:"progress_uepoch"`
		ProgressMediaEpoch          string    `xml:"progress_media_epoch"`
		ProgressMediaUepoch         string    `xml:"progress_media_uepoch"`
		EndEpoch                    xmlInt64  `xml:"end_epoch"`
		EndUepoch                   string    `xml:"end_uepoch"`
		LastApp                     string    `xml:"last_app"`
		LastArg                     string    `xml:"last_arg"`
		CallerID                    string    `xml:"caller_id"`
		Duration                    xmlInt64  `xml:"duration"`
		Billsec                     xmlInt64  `xml:"billsec"`
		Progresssec                 string    `xml:"progresssec"`
		Answersec                   string    `xml:"answersec"`
		Waitsec                     string    `xml:"waitsec"`
		ProgressMediasec            string    `xml:"progress_mediasec"`
		FlowBillsec                 string    `xml:"flow_billsec"`
		Mduration                   string    `xml:"mduration"`
		Billmsec                    string    `xml:"billmsec"`
		Progressmsec                string    `xml:"progressmsec"`
		Answermsec                  string    `xml:"answermsec"`
		Waitmsec                    string    `xml:"waitmsec"`
		ProgressMediamsec           string    `xml:"progress_mediamsec"`
		FlowBillmsec                string    `xml:"flow_billmsec"`
		Uduration                   string    `xml:"uduration"`
		Billusec                    string    `xml:"billusec"`
		Progressusec                string    `xml:"progressusec"`
		Answerusec                  string    `xml:"answerusec"`
		Waitusec                    string    `xml:"waitusec"`
		ProgressMediausec           string    `xml:"progress_mediausec"`
		FlowBillusec                string    `xml:"flow_billusec"`
		RTPAudioInRawBytes          string    `xml:"rtp_audio_in_raw_bytes"`
		RTPAudioInMediaBytes        string    `xml:"rtp_audio_in_media_bytes"`
		RTPAudioInPacketCount       string    `xml:"rtp_audio_in_packet_count"`
		RTPAudioInMediaPacketCount  string    `xml:"rtp_audio_in_media_packet_count"`
		RTPAudioInSkipPacketCount   string    `xml:"rtp_audio_in_skip_packet_count"`
		RTPAudioInJitterPacketCount string    `xml:"rtp_audio_in_jitter_packet_count"`
		RTPAudioInDtmfPacketCount   string    `xml:"rtp_audio_in_dtmf_packet_count"`
		RTPAudioInCngPacketCount    string    `xml:"rtp_audio_in_cng_packet_count"`
		RTPAudioInFlushPacketCount  string    `xml:"rtp_audio_in_flush_packet_count"`
		RTPAudioInLargestJbSize     string    `xml:"rtp_audio_in_largest_jb_size"`
		RTPAudioInJitterMinVariance string    `xml:"rtp_audio_in_jitter_min_variance"`
		RTPAudioInJitterMaxVariance string    `xml:"rtp_audio_in_jitter_max_variance"`
		RTPAudioInJitterLossRate    string    `xml:"rtp_audio_in_jitter_loss_rate"`
		RTPAudioInJitterBurstRate   string    `xml:"rtp_audio_in_jitter_burst_rate"`
		RTPAudioInMeanInterval      string    `xml:"rtp_audio_in_mean_interval"`
		RTPAudioInFlawTotal         string    `xml:"rtp_audio_in_flaw_total"`
		RTPAudioInQualityPercentage string    `xml:"rtp_audio_in_quality_percentage"`
		RTPAudioInMos               string    `xml:"rtp_audio_in_mos"`
		RTPAudioOutRawBytes         string    `xml:"rtp_audio_out_raw_bytes"`
		RTPAudioOutMediaBytes       string    `xml:"rtp_audio_out_media_bytes"`
		RTPAudioOutPacketCount      string    `xml:"rtp_audio_out_packet_count"`
		RTPAudioOutMediaPacketCount string    `xml:"rtp_audio_out_media_packet_count"`
		RTPAudioOutSkipPacketCount  string    `xml:"rtp_audio_out_skip_packet_count"`
		RTPAudioOutDtmfPacketCount  string    `xml:"rtp_audio_out_dtmf_packet_count"`
		RTPAudioOutCngPacketCount   string    `xml:"rtp_audio_out_cng_packet_count"`
		RTPAudioRtcpPacketCount     string    `xml:"rtp_audio_rtcp_packet_count"`
		RTPAudioRtcpOctetCount      string    `xml:"rtp_audio_rtcp_octet_count"`
		CJCdr                       string    `xml:"cj_cdr"`
		CJCallDirection             string    `xml:"cj_direction"`
		CJRecordName                xmlString `xml:"cj_record_name"`
		CJRecordPrtTag              string    `xml:"cj_record_prt_tag"`
		CJInfo1                     string    `xml:"cj_info1"`
		CJInfo2                     string    `xml:"cj_info2"`
		CJInfo3                     string    `xml:"cj_info3"`
		CJInfo4                     string    `xml:"cj_info4"`
		CJInfo5                     string    `xml:"cj_info5"`
		CJInfo6                     string    `xml:"cj_info6"`
		CJInfo7                     string    `xml:"cj_info7"`
		CJInfo8                     string    `xml:"cj_info8"`
		CJInfo9                     string    `xml:"cj_info9"`
		CJInfo10                    string    `xml:"cj_info10"`
	} `xml:"variables"`
	AppLog struct {
		Text        string `xml:",chardata"`
		Application []struct {
			Text     string `xml:",chardata"`
			AppName  string `xml:"app_name,attr"`
			AppData  string `xml:"app_data,attr"`
			AppStamp string `xml:"app_stamp,attr"`
		} `xml:"application"`
	} `xml:"app_log"`
	Callflow []struct {
		Text         string `xml:",chardata"`
		Dialplan     string `xml:"dialplan,attr"`
		UniqueID     string `xml:"unique-id,attr"`
		CloneOf      string `xml:"clone-of,attr"`
		ProfileIndex string `xml:"profile_index,attr"`
		Extension    struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name,attr"`
			Number      string `xml:"number,attr"`
			CurrentApp  string `xml:"current_app,attr"`
			Application []struct {
				Text         string `xml:",chardata"`
				AppName      string `xml:"app_name,attr"`
				AppData      string `xml:"app_data,attr"`
				LastExecuted string `xml:"last_executed,attr"`
			} `xml:"application"`
		} `xml:"extension"`
		CallerProfile struct {
			Text              string    `xml:",chardata"`
			Username          xmlString `xml:"username"`
			Dialplan          string    `xml:"dialplan"`
			CallerIDName      xmlString `xml:"caller_id_name"`
			CallerIDNumber    xmlString `xml:"caller_id_number"`
			CalleeIDName      xmlString `xml:"callee_id_name"`
			CalleeIDNumber    xmlString `xml:"callee_id_number"`
			Ani               string    `xml:"ani"`
			Aniii             string    `xml:"aniii"`
			NetworkAddr       string    `xml:"network_addr"`
			Rdnis             string    `xml:"rdnis"`
			DestinationNumber xmlString `xml:"destination_number"`
			UUID              string    `xml:"uuid"`
			Source            string    `xml:"source"`
			TransferSource    string    `xml:"transfer_source"`
			Context           string    `xml:"context"`
			ChanName          string    `xml:"chan_name"`
			Origination       struct {
				Text                     string `xml:",chardata"`
				OriginationCallerProfile struct {
					Text              string    `xml:",chardata"`
					Username          xmlString `xml:"username"`
					Dialplan          string    `xml:"dialplan"`
					CallerIDName      xmlString `xml:"caller_id_name"`
					CallerIDNumber    xmlString `xml:"caller_id_number"`
					CalleeIDName      xmlString `xml:"callee_id_name"`
					CalleeIDNumber    xmlString `xml:"callee_id_number"`
					Ani               string    `xml:"ani"`
					Aniii             string    `xml:"aniii"`
					NetworkAddr       string    `xml:"network_addr"`
					Rdnis             string    `xml:"rdnis"`
					DestinationNumber xmlString `xml:"destination_number"`
					UUID              string    `xml:"uuid"`
					Source            string    `xml:"source"`
					Context           string    `xml:"context"`
					ChanName          string    `xml:"chan_name"`
				} `xml:"origination_caller_profile"`
			} `xml:"origination"`
			Originatee struct {
				Text                    string `xml:",chardata"`
				OriginateeCallerProfile struct {
					Text              string    `xml:",chardata"`
					Username          xmlString `xml:"username"`
					Dialplan          string    `xml:"dialplan"`
					CallerIDName      xmlString `xml:"caller_id_name"`
					CallerIDNumber    xmlString `xml:"caller_id_number"`
					CalleeIDName      xmlString `xml:"callee_id_name"`
					CalleeIDNumber    xmlString `xml:"callee_id_number"`
					Ani               string    `xml:"ani"`
					Aniii             string    `xml:"aniii"`
					NetworkAddr       string    `xml:"network_addr"`
					Rdnis             string    `xml:"rdnis"`
					DestinationNumber xmlString `xml:"destination_number"`
					UUID              string    `xml:"uuid"`
					Source            string    `xml:"source"`
					Context           string    `xml:"context"`
					ChanName          string    `xml:"chan_name"`
				} `xml:"originatee_caller_profile"`
			} `xml:"originatee"`
		} `xml:"caller_profile"`
		Times struct {
			Text               string `xml:",chardata"`
			CreatedTime        string `xml:"created_time"`
			ProfileCreatedTime string `xml:"profile_created_time"`
			ProgressTime       string `xml:"progress_time"`
			ProgressMediaTime  string `xml:"progress_media_time"`
			AnsweredTime       string `xml:"answered_time"`
			BridgedTime        string `xml:"bridged_time"`
			LastHoldTime       string `xml:"last_hold_time"`
			HoldAccumTime      string `xml:"hold_accum_time"`
			HangupTime         string `xml:"hangup_time"`
			ResurrectTime      string `xml:"resurrect_time"`
			TransferTime       string `xml:"transfer_time"`
		} `xml:"times"`
	} `xml:"callflow"`
}

type xmlString string

func (u *xmlString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	str, err := url.QueryUnescape(content)
	if err != nil {
		return err
	}

	*u = xmlString(str)

	return nil
}

type xmlInt64 int64

func (u *xmlInt64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	n, err := strconv.Atoi(content)
	if err != nil {
		return err
	}

	*u = xmlInt64(n)

	return nil
}
