// api/core.ts
var NEUTRON = {};

// api/logger.ts
var LogLevel;
((LogLevel2) => {
  LogLevel2[LogLevel2["DEBUG"] = -4] = "DEBUG";
  LogLevel2[LogLevel2["INFO"] = 0] = "INFO";
  LogLevel2[LogLevel2["WARN"] = 4] = "WARN";
  LogLevel2[LogLevel2["ERROR"] = 8] = "ERROR";
  LogLevel2[LogLevel2["FATAL"] = 12] = "FATAL";
})(LogLevel ||= {});

// api/index.ts
Object.defineProperty(NEUTRON, "LogLevel", LogLevel);
Object.defineProperty(window, "NEUTRON", NEUTRON);

//# debugId=A450C8AFA9EE878864756E2164756E21
//# sourceMappingURL=index.js.map
