// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./channelAlexa";
export * from "./channelDirectLine";
export * from "./channelDirectLineSpeech";
export * from "./channelEmail";
export * from "./channelFacebook";
export * from "./channelLine";
export * from "./channelSlack";
export * from "./channelSms";
export * from "./channelTeams";
export * from "./channelWebChat";
export * from "./channelsRegistration";
export * from "./connection";
export * from "./serviceAzureBot";
export * from "./webApp";

// Import resources to register:
import { ChannelAlexa } from "./channelAlexa";
import { ChannelDirectLine } from "./channelDirectLine";
import { ChannelDirectLineSpeech } from "./channelDirectLineSpeech";
import { ChannelEmail } from "./channelEmail";
import { ChannelFacebook } from "./channelFacebook";
import { ChannelLine } from "./channelLine";
import { ChannelSlack } from "./channelSlack";
import { ChannelSms } from "./channelSms";
import { ChannelTeams } from "./channelTeams";
import { ChannelWebChat } from "./channelWebChat";
import { ChannelsRegistration } from "./channelsRegistration";
import { Connection } from "./connection";
import { ServiceAzureBot } from "./serviceAzureBot";
import { WebApp } from "./webApp";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:bot/channelAlexa:ChannelAlexa":
                return new ChannelAlexa(name, <any>undefined, { urn })
            case "azure:bot/channelDirectLine:ChannelDirectLine":
                return new ChannelDirectLine(name, <any>undefined, { urn })
            case "azure:bot/channelDirectLineSpeech:ChannelDirectLineSpeech":
                return new ChannelDirectLineSpeech(name, <any>undefined, { urn })
            case "azure:bot/channelEmail:ChannelEmail":
                return new ChannelEmail(name, <any>undefined, { urn })
            case "azure:bot/channelFacebook:ChannelFacebook":
                return new ChannelFacebook(name, <any>undefined, { urn })
            case "azure:bot/channelLine:ChannelLine":
                return new ChannelLine(name, <any>undefined, { urn })
            case "azure:bot/channelSlack:ChannelSlack":
                return new ChannelSlack(name, <any>undefined, { urn })
            case "azure:bot/channelSms:ChannelSms":
                return new ChannelSms(name, <any>undefined, { urn })
            case "azure:bot/channelTeams:ChannelTeams":
                return new ChannelTeams(name, <any>undefined, { urn })
            case "azure:bot/channelWebChat:ChannelWebChat":
                return new ChannelWebChat(name, <any>undefined, { urn })
            case "azure:bot/channelsRegistration:ChannelsRegistration":
                return new ChannelsRegistration(name, <any>undefined, { urn })
            case "azure:bot/connection:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "azure:bot/serviceAzureBot:ServiceAzureBot":
                return new ServiceAzureBot(name, <any>undefined, { urn })
            case "azure:bot/webApp:WebApp":
                return new WebApp(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "bot/channelAlexa", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelDirectLine", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelDirectLineSpeech", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelEmail", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelFacebook", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelLine", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelSlack", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelSms", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelTeams", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelWebChat", _module)
pulumi.runtime.registerResourceModule("azure", "bot/channelsRegistration", _module)
pulumi.runtime.registerResourceModule("azure", "bot/connection", _module)
pulumi.runtime.registerResourceModule("azure", "bot/serviceAzureBot", _module)
pulumi.runtime.registerResourceModule("azure", "bot/webApp", _module)
