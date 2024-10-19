import { ObjectFlags } from "typescript";
import { NEUTRON } from "./core";
import { LogLevel } from "./logger";

Object.defineProperty(NEUTRON, "LogLevel", LogLevel);
Object.defineProperty(window, "NEUTRON", NEUTRON);
