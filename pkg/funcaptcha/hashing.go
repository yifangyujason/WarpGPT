// https://github.com/fingerprintjs/fingerprintjs/blob/master/src/utils/hashing.ts

package funcaptcha

import (
	"fmt"
	"sort"
	"strings"
)

//goland:noinspection SpellCheckingInspection
func getWindowHash() string { // return aA(b1[df(f_a_hT.e)]()[df(f_a_hT.f)]('|'), 0x1a4);
	// Object.getOwnPropertyNames(window);
	b1 := []string{
		"ALFCCJS",
		"AbortController",
		"AbortSignal",
		"AbsoluteOrientationSensor",
		"AbstractRange",
		"Accelerometer",
		"AggregateError",
		"AnalyserNode",
		"Animation",
		"AnimationEffect",
		"AnimationEvent",
		"AnimationPlaybackEvent",
		"AnimationTimeline",
		"ArkoseEnforcement",
		"Array",
		"ArrayBuffer",
		"Atomics",
		"Attr",
		"Audio",
		"AudioBuffer",
		"AudioBufferSourceNode",
		"AudioContext",
		"AudioData",
		"AudioDecoder",
		"AudioDestinationNode",
		"AudioEncoder",
		"AudioListener",
		"AudioNode",
		"AudioParam",
		"AudioParamMap",
		"AudioProcessingEvent",
		"AudioScheduledSourceNode",
		"AudioSinkInfo",
		"AudioWorklet",
		"AudioWorkletNode",
		"AuthenticatorAssertionResponse",
		"AuthenticatorAttestationResponse",
		"AuthenticatorResponse",
		"BackgroundFetchManager",
		"BackgroundFetchRecord",
		"BackgroundFetchRegistration",
		"BarProp",
		"BaseAudioContext",
		"BatteryManager",
		"BeforeInstallPromptEvent",
		"BeforeUnloadEvent",
		"BigInt",
		"BigInt64Array",
		"BigUint64Array",
		"BiquadFilterNode",
		"Blob",
		"BlobEvent",
		"Boolean",
		"BroadcastChannel",
		"BrowserCaptureMediaStreamTrack",
		"ByteLengthQueuingStrategy",
		"CDATASection",
		"CSS",
		"CSSAnimation",
		"CSSConditionRule",
		"CSSContainerRule",
		"CSSCounterStyleRule",
		"CSSFontFaceRule",
		"CSSFontPaletteValuesRule",
		"CSSGroupingRule",
		"CSSImageValue",
		"CSSImportRule",
		"CSSKeyframeRule",
		"CSSKeyframesRule",
		"CSSKeywordValue",
		"CSSLayerBlockRule",
		"CSSLayerStatementRule",
		"CSSMathClamp",
		"CSSMathInvert",
		"CSSMathMax",
		"CSSMathMin",
		"CSSMathNegate",
		"CSSMathProduct",
		"CSSMathSum",
		"CSSMathValue",
		"CSSMatrixComponent",
		"CSSMediaRule",
		"CSSNamespaceRule",
		"CSSNumericArray",
		"CSSNumericValue",
		"CSSPageRule",
		"CSSPerspective",
		"CSSPositionValue",
		"CSSPropertyRule",
		"CSSRotate",
		"CSSRule",
		"CSSRuleList",
		"CSSScale",
		"CSSSkew",
		"CSSSkewX",
		"CSSSkewY",
		"CSSStyleDeclaration",
		"CSSStyleRule",
		"CSSStyleSheet",
		"CSSStyleValue",
		"CSSSupportsRule",
		"CSSTransformComponent",
		"CSSTransformValue",
		"CSSTransition",
		"CSSTranslate",
		"CSSUnitValue",
		"CSSUnparsedValue",
		"CSSVariableReferenceValue",
		"Cache",
		"CacheStorage",
		"CanvasCaptureMediaStreamTrack",
		"CanvasGradient",
		"CanvasPattern",
		"CanvasRenderingContext2D",
		"CaptureController",
		"ChannelMergerNode",
		"ChannelSplitterNode",
		"CharacterData",
		"Clipboard",
		"ClipboardEvent",
		"ClipboardItem",
		"CloseEvent",
		"Comment",
		"CompositionEvent",
		"CompressionStream",
		"ConstantSourceNode",
		"ContentVisibilityAutoStateChangeEvent",
		"ConvolverNode",
		"CookieChangeEvent",
		"CookieStore",
		"CookieStoreManager",
		"CountQueuingStrategy",
		"Credential",
		"CredentialsContainer",
		"CropTarget",
		"Crypto",
		"CryptoKey",
		"CustomElementRegistry",
		"CustomEvent",
		"CustomStateSet",
		"DOMError",
		"DOMException",
		"DOMImplementation",
		"DOMMatrix",
		"DOMMatrixReadOnly",
		"DOMParser",
		"DOMPoint",
		"DOMPointReadOnly",
		"DOMQuad",
		"DOMRect",
		"DOMRectList",
		"DOMRectReadOnly",
		"DOMStringList",
		"DOMStringMap",
		"DOMTokenList",
		"DataTransfer",
		"DataTransferItem",
		"DataTransferItemList",
		"DataView",
		"Date",
		"DecompressionStream",
		"DelayNode",
		"DelegatedInkTrailPresenter",
		"DeviceMotionEvent",
		"DeviceMotionEventAcceleration",
		"DeviceMotionEventRotationRate",
		"DeviceOrientationEvent",
		"Document",
		"DocumentFragment",
		"DocumentTimeline",
		"DocumentType",
		"DragEvent",
		"DynamicsCompressorNode",
		"Element",
		"ElementInternals",
		"EncodedAudioChunk",
		"EncodedVideoChunk",
		"Error",
		"ErrorEvent",
		"EvalError",
		"Event",
		"EventCounts",
		"EventSource",
		"EventTarget",
		"External",
		"EyeDropper",
		"FeaturePolicy",
		"FederatedCredential",
		"File",
		"FileList",
		"FileReader",
		"FileSystemDirectoryHandle",
		"FileSystemFileHandle",
		"FileSystemHandle",
		"FileSystemWritableFileStream",
		"FinalizationRegistry",
		"Float32Array",
		"Float64Array",
		"FocusEvent",
		"FontData",
		"FontFace",
		"FontFaceSetLoadEvent",
		"FormData",
		"FormDataEvent",
		"FragmentDirective",
		"FunCaptcha",
		"Function",
		"GPU",
		"GPUAdapter",
		"GPUAdapterInfo",
		"GPUBindGroup",
		"GPUBindGroupLayout",
		"GPUBuffer",
		"GPUBufferUsage",
		"GPUCanvasContext",
		"GPUColorWrite",
		"GPUCommandBuffer",
		"GPUCommandEncoder",
		"GPUCompilationInfo",
		"GPUCompilationMessage",
		"GPUComputePassEncoder",
		"GPUComputePipeline",
		"GPUDevice",
		"GPUDeviceLostInfo",
		"GPUError",
		"GPUExternalTexture",
		"GPUInternalError",
		"GPUMapMode",
		"GPUOutOfMemoryError",
		"GPUPipelineError",
		"GPUPipelineLayout",
		"GPUQuerySet",
		"GPUQueue",
		"GPURenderBundle",
		"GPURenderBundleEncoder",
		"GPURenderPassEncoder",
		"GPURenderPipeline",
		"GPUSampler",
		"GPUShaderModule",
		"GPUShaderStage",
		"GPUSupportedFeatures",
		"GPUSupportedLimits",
		"GPUTexture",
		"GPUTextureUsage",
		"GPUTextureView",
		"GPUUncapturedErrorEvent",
		"GPUValidationError",
		"GainNode",
		"Gamepad",
		"GamepadButton",
		"GamepadEvent",
		"GamepadHapticActuator",
		"Geolocation",
		"GeolocationCoordinates",
		"GeolocationPosition",
		"GeolocationPositionError",
		"GravitySensor",
		"Gyroscope",
		"HID",
		"HIDConnectionEvent",
		"HIDDevice",
		"HIDInputReportEvent",
		"HTMLAllCollection",
		"HTMLAnchorElement",
		"HTMLAreaElement",
		"HTMLAudioElement",
		"HTMLBRElement",
		"HTMLBaseElement",
		"HTMLBodyElement",
		"HTMLButtonElement",
		"HTMLCanvasElement",
		"HTMLCollection",
		"HTMLDListElement",
		"HTMLDataElement",
		"HTMLDataListElement",
		"HTMLDetailsElement",
		"HTMLDialogElement",
		"HTMLDirectoryElement",
		"HTMLDivElement",
		"HTMLDocument",
		"HTMLElement",
		"HTMLEmbedElement",
		"HTMLFieldSetElement",
		"HTMLFontElement",
		"HTMLFormControlsCollection",
		"HTMLFormElement",
		"HTMLFrameElement",
		"HTMLFrameSetElement",
		"HTMLHRElement",
		"HTMLHeadElement",
		"HTMLHeadingElement",
		"HTMLHtmlElement",
		"HTMLIFrameElement",
		"HTMLImageElement",
		"HTMLInputElement",
		"HTMLLIElement",
		"HTMLLabelElement",
		"HTMLLegendElement",
		"HTMLLinkElement",
		"HTMLMapElement",
		"HTMLMarqueeElement",
		"HTMLMediaElement",
		"HTMLMenuElement",
		"HTMLMetaElement",
		"HTMLMeterElement",
		"HTMLModElement",
		"HTMLOListElement",
		"HTMLObjectElement",
		"HTMLOptGroupElement",
		"HTMLOptionElement",
		"HTMLOptionsCollection",
		"HTMLOutputElement",
		"HTMLParagraphElement",
		"HTMLParamElement",
		"HTMLPictureElement",
		"HTMLPreElement",
		"HTMLProgressElement",
		"HTMLQuoteElement",
		"HTMLScriptElement",
		"HTMLSelectElement",
		"HTMLSlotElement",
		"HTMLSourceElement",
		"HTMLSpanElement",
		"HTMLStyleElement",
		"HTMLTableCaptionElement",
		"HTMLTableCellElement",
		"HTMLTableColElement",
		"HTMLTableElement",
		"HTMLTableRowElement",
		"HTMLTableSectionElement",
		"HTMLTemplateElement",
		"HTMLTextAreaElement",
		"HTMLTimeElement",
		"HTMLTitleElement",
		"HTMLTrackElement",
		"HTMLUListElement",
		"HTMLUnknownElement",
		"HTMLVideoElement",
		"HashChangeEvent",
		"Headers",
		"Highlight",
		"HighlightRegistry",
		"History",
		"IDBCursor",
		"IDBCursorWithValue",
		"IDBDatabase",
		"IDBFactory",
		"IDBIndex",
		"IDBKeyRange",
		"IDBObjectStore",
		"IDBOpenDBRequest",
		"IDBRequest",
		"IDBTransaction",
		"IDBVersionChangeEvent",
		"IIRFilterNode",
		"IdentityCredential",
		"IdleDeadline",
		"IdleDetector",
		"Image",
		"ImageBitmap",
		"ImageBitmapRenderingContext",
		"ImageCapture",
		"ImageData",
		"ImageDecoder",
		"ImageTrack",
		"ImageTrackList",
		"Infinity",
		"Ink",
		"InputDeviceCapabilities",
		"InputDeviceInfo",
		"InputEvent",
		"Int16Array",
		"Int32Array",
		"Int8Array",
		"IntersectionObserver",
		"IntersectionObserverEntry",
		"Intl",
		"JSON",
		"Keyboard",
		"KeyboardEvent",
		"KeyboardLayoutMap",
		"KeyframeEffect",
		"LargestContentfulPaint",
		"LaunchParams",
		"LaunchQueue",
		"LayoutShift",
		"LayoutShiftAttribution",
		"LinearAccelerationSensor",
		"Location",
		"Lock",
		"LockManager",
		"MIDIAccess",
		"MIDIConnectionEvent",
		"MIDIInput",
		"MIDIInputMap",
		"MIDIMessageEvent",
		"MIDIOutput",
		"MIDIOutputMap",
		"MIDIPort",
		"Map",
		"Math",
		"MathMLElement",
		"MediaCapabilities",
		"MediaDeviceInfo",
		"MediaDevices",
		"MediaElementAudioSourceNode",
		"MediaEncryptedEvent",
		"MediaError",
		"MediaKeyMessageEvent",
		"MediaKeySession",
		"MediaKeyStatusMap",
		"MediaKeySystemAccess",
		"MediaKeys",
		"MediaList",
		"MediaMetadata",
		"MediaQueryList",
		"MediaQueryListEvent",
		"MediaRecorder",
		"MediaSession",
		"MediaSource",
		"MediaSourceHandle",
		"MediaStream",
		"MediaStreamAudioDestinationNode",
		"MediaStreamAudioSourceNode",
		"MediaStreamEvent",
		"MediaStreamTrack",
		"MediaStreamTrackEvent",
		"MediaStreamTrackGenerator",
		"MediaStreamTrackProcessor",
		"MessageChannel",
		"MessageEvent",
		"MessagePort",
		"MimeType",
		"MimeTypeArray",
		"MouseEvent",
		"MutationEvent",
		"MutationObserver",
		"MutationRecord",
		"NaN",
		"NamedNodeMap",
		"NavigateEvent",
		"Navigation",
		"NavigationCurrentEntryChangeEvent",
		"NavigationDestination",
		"NavigationHistoryEntry",
		"NavigationPreloadManager",
		"NavigationTransition",
		"Navigator",
		"NavigatorManagedData",
		"NavigatorUAData",
		"NetworkInformation",
		"Node",
		"NodeFilter",
		"NodeIterator",
		"NodeList",
		"Notification",
		"Number",
		"OTPCredential",
		"Object",
		"OfflineAudioCompletionEvent",
		"OfflineAudioContext",
		"OffscreenCanvas",
		"OffscreenCanvasRenderingContext2D",
		"Option",
		"OrientationSensor",
		"OscillatorNode",
		"OverconstrainedError",
		"PageTransitionEvent",
		"PannerNode",
		"PasswordCredential",
		"Path2D",
		"PaymentAddress",
		"PaymentManager",
		"PaymentMethodChangeEvent",
		"PaymentRequest",
		"PaymentRequestUpdateEvent",
		"PaymentResponse",
		"Performance",
		"PerformanceElementTiming",
		"PerformanceEntry",
		"PerformanceEventTiming",
		"PerformanceLongTaskTiming",
		"PerformanceMark",
		"PerformanceMeasure",
		"PerformanceNavigation",
		"PerformanceNavigationTiming",
		"PerformanceObserver",
		"PerformanceObserverEntryList",
		"PerformancePaintTiming",
		"PerformanceResourceTiming",
		"PerformanceServerTiming",
		"PerformanceTiming",
		"PeriodicSyncManager",
		"PeriodicWave",
		"PermissionStatus",
		"Permissions",
		"PictureInPictureEvent",
		"PictureInPictureWindow",
		"Plugin",
		"PluginArray",
		"PointerEvent",
		"PopStateEvent",
		"Presentation",
		"PresentationAvailability",
		"PresentationConnection",
		"PresentationConnectionAvailableEvent",
		"PresentationConnectionCloseEvent",
		"PresentationConnectionList",
		"PresentationReceiver",
		"PresentationRequest",
		"ProcessingInstruction",
		"Profiler",
		"ProgressEvent",
		"Promise",
		"PromiseRejectionEvent",
		"Proxy",
		"PublicKeyCredential",
		"PushManager",
		"PushSubscription",
		"PushSubscriptionOptions",
		"RTCCertificate",
		"RTCDTMFSender",
		"RTCDTMFToneChangeEvent",
		"RTCDataChannel",
		"RTCDataChannelEvent",
		"RTCDtlsTransport",
		"RTCEncodedAudioFrame",
		"RTCEncodedVideoFrame",
		"RTCError",
		"RTCErrorEvent",
		"RTCIceCandidate",
		"RTCIceTransport",
		"RTCPeerConnection",
		"RTCPeerConnectionIceErrorEvent",
		"RTCPeerConnectionIceEvent",
		"RTCRtpReceiver",
		"RTCRtpSender",
		"RTCRtpTransceiver",
		"RTCSctpTransport",
		"RTCSessionDescription",
		"RTCStatsReport",
		"RTCTrackEvent",
		"RadioNodeList",
		"Range",
		"RangeError",
		"ReadableByteStreamController",
		"ReadableStream",
		"ReadableStreamBYOBReader",
		"ReadableStreamBYOBRequest",
		"ReadableStreamDefaultController",
		"ReadableStreamDefaultReader",
		"ReferenceError",
		"Reflect",
		"RegExp",
		"RelativeOrientationSensor",
		"RemotePlayback",
		"ReportingObserver",
		"Request",
		"ResizeObserver",
		"ResizeObserverEntry",
		"ResizeObserverSize",
		"Response",
		"SVGAElement",
		"SVGAngle",
		"SVGAnimateElement",
		"SVGAnimateMotionElement",
		"SVGAnimateTransformElement",
		"SVGAnimatedAngle",
		"SVGAnimatedBoolean",
		"SVGAnimatedEnumeration",
		"SVGAnimatedInteger",
		"SVGAnimatedLength",
		"SVGAnimatedLengthList",
		"SVGAnimatedNumber",
		"SVGAnimatedNumberList",
		"SVGAnimatedPreserveAspectRatio",
		"SVGAnimatedRect",
		"SVGAnimatedString",
		"SVGAnimatedTransformList",
		"SVGAnimationElement",
		"SVGCircleElement",
		"SVGClipPathElement",
		"SVGComponentTransferFunctionElement",
		"SVGDefsElement",
		"SVGDescElement",
		"SVGElement",
		"SVGEllipseElement",
		"SVGFEBlendElement",
		"SVGFEColorMatrixElement",
		"SVGFEComponentTransferElement",
		"SVGFECompositeElement",
		"SVGFEConvolveMatrixElement",
		"SVGFEDiffuseLightingElement",
		"SVGFEDisplacementMapElement",
		"SVGFEDistantLightElement",
		"SVGFEDropShadowElement",
		"SVGFEFloodElement",
		"SVGFEFuncAElement",
		"SVGFEFuncBElement",
		"SVGFEFuncGElement",
		"SVGFEFuncRElement",
		"SVGFEGaussianBlurElement",
		"SVGFEImageElement",
		"SVGFEMergeElement",
		"SVGFEMergeNodeElement",
		"SVGFEMorphologyElement",
		"SVGFEOffsetElement",
		"SVGFEPointLightElement",
		"SVGFESpecularLightingElement",
		"SVGFESpotLightElement",
		"SVGFETileElement",
		"SVGFETurbulenceElement",
		"SVGFilterElement",
		"SVGForeignObjectElement",
		"SVGGElement",
		"SVGGeometryElement",
		"SVGGradientElement",
		"SVGGraphicsElement",
		"SVGImageElement",
		"SVGLength",
		"SVGLengthList",
		"SVGLineElement",
		"SVGLinearGradientElement",
		"SVGMPathElement",
		"SVGMarkerElement",
		"SVGMaskElement",
		"SVGMatrix",
		"SVGMetadataElement",
		"SVGNumber",
		"SVGNumberList",
		"SVGPathElement",
		"SVGPatternElement",
		"SVGPoint",
		"SVGPointList",
		"SVGPolygonElement",
		"SVGPolylineElement",
		"SVGPreserveAspectRatio",
		"SVGRadialGradientElement",
		"SVGRect",
		"SVGRectElement",
		"SVGSVGElement",
		"SVGScriptElement",
		"SVGSetElement",
		"SVGStopElement",
		"SVGStringList",
		"SVGStyleElement",
		"SVGSwitchElement",
		"SVGSymbolElement",
		"SVGTSpanElement",
		"SVGTextContentElement",
		"SVGTextElement",
		"SVGTextPathElement",
		"SVGTextPositioningElement",
		"SVGTitleElement",
		"SVGTransform",
		"SVGTransformList",
		"SVGUnitTypes",
		"SVGUseElement",
		"SVGViewElement",
		"Sanitizer",
		"Scheduler",
		"Scheduling",
		"Screen",
		"ScreenDetailed",
		"ScreenDetails",
		"ScreenOrientation",
		"ScriptProcessorNode",
		"SecurityPolicyViolationEvent",
		"Selection",
		"Sensor",
		"SensorErrorEvent",
		"Serial",
		"SerialPort",
		"ServiceWorker",
		"ServiceWorkerContainer",
		"ServiceWorkerRegistration",
		"Set",
		"ShadowRoot",
		"SharedWorker",
		"SourceBuffer",
		"SourceBufferList",
		"SpeechSynthesisErrorEvent",
		"SpeechSynthesisEvent",
		"SpeechSynthesisUtterance",
		"StaticRange",
		"StereoPannerNode",
		"Storage",
		"StorageEvent",
		"StorageManager",
		"String",
		"StylePropertyMap",
		"StylePropertyMapReadOnly",
		"StyleSheet",
		"StyleSheetList",
		"SubmitEvent",
		"SubtleCrypto",
		"Symbol",
		"SyncManager",
		"SyntaxError",
		"TaskAttributionTiming",
		"TaskController",
		"TaskPriorityChangeEvent",
		"TaskSignal",
		"Text",
		"TextDecoder",
		"TextDecoderStream",
		"TextEncoder",
		"TextEncoderStream",
		"TextEvent",
		"TextMetrics",
		"TextTrack",
		"TextTrackCue",
		"TextTrackCueList",
		"TextTrackList",
		"TimeRanges",
		"ToggleEvent",
		"Touch",
		"TouchEvent",
		"TouchList",
		"TrackEvent",
		"TransformStream",
		"TransformStreamDefaultController",
		"TransitionEvent",
		"TreeWalker",
		"TrustedHTML",
		"TrustedScript",
		"TrustedScriptURL",
		"TrustedTypePolicy",
		"TrustedTypePolicyFactory",
		"TypeError",
		"UIEvent",
		"URIError",
		"URL",
		"URLPattern",
		"URLSearchParams",
		"USB",
		"USBAlternateInterface",
		"USBConfiguration",
		"USBConnectionEvent",
		"USBDevice",
		"USBEndpoint",
		"USBInTransferResult",
		"USBInterface",
		"USBIsochronousInTransferPacket",
		"USBIsochronousInTransferResult",
		"USBIsochronousOutTransferPacket",
		"USBIsochronousOutTransferResult",
		"USBOutTransferResult",
		"Uint16Array",
		"Uint32Array",
		"Uint8Array",
		"Uint8ClampedArray",
		"UserActivation",
		"VTTCue",
		"ValidityState",
		"VideoColorSpace",
		"VideoDecoder",
		"VideoEncoder",
		"VideoFrame",
		"VideoPlaybackQuality",
		"ViewTransition",
		"VirtualKeyboard",
		"VirtualKeyboardGeometryChangeEvent",
		"VisualViewport",
		"WakeLock",
		"WakeLockSentinel",
		"WaveShaperNode",
		"WeakMap",
		"WeakRef",
		"WeakSet",
		"WebAssembly",
		"WebGL2RenderingContext",
		"WebGLActiveInfo",
		"WebGLBuffer",
		"WebGLContextEvent",
		"WebGLFramebuffer",
		"WebGLProgram",
		"WebGLQuery",
		"WebGLRenderbuffer",
		"WebGLRenderingContext",
		"WebGLSampler",
		"WebGLShader",
		"WebGLShaderPrecisionFormat",
		"WebGLSync",
		"WebGLTexture",
		"WebGLTransformFeedback",
		"WebGLUniformLocation",
		"WebGLVertexArrayObject",
		"WebKitCSSMatrix",
		"WebKitMutationObserver",
		"WebSocket",
		"WebTransport",
		"WebTransportBidirectionalStream",
		"WebTransportDatagramDuplexStream",
		"WebTransportError",
		"WheelEvent",
		"Window",
		"WindowControlsOverlay",
		"WindowControlsOverlayGeometryChangeEvent",
		"Worker",
		"Worklet",
		"WritableStream",
		"WritableStreamDefaultController",
		"WritableStreamDefaultWriter",
		"XMLDocument",
		"XMLHttpRequest",
		"XMLHttpRequestEventTarget",
		"XMLHttpRequestUpload",
		"XMLSerializer",
		"XPathEvaluator",
		"XPathExpression",
		"XPathResult",
		"XRAnchor",
		"XRAnchorSet",
		"XRBoundedReferenceSpace",
		"XRCPUDepthInformation",
		"XRCamera",
		"XRDOMOverlayState",
		"XRDepthInformation",
		"XRFrame",
		"XRHitTestResult",
		"XRHitTestSource",
		"XRInputSource",
		"XRInputSourceArray",
		"XRInputSourceEvent",
		"XRInputSourcesChangeEvent",
		"XRLayer",
		"XRLightEstimate",
		"XRLightProbe",
		"XRPose",
		"XRRay",
		"XRReferenceSpace",
		"XRReferenceSpaceEvent",
		"XRRenderState",
		"XRRigidTransform",
		"XRSession",
		"XRSessionEvent",
		"XRSpace",
		"XRSystem",
		"XRTransientInputHitTestResult",
		"XRTransientInputHitTestSource",
		"XRView",
		"XRViewerPose",
		"XRViewport",
		"XRWebGLBinding",
		"XRWebGLDepthInformation",
		"XRWebGLLayer",
		"XSLTProcessor",
		"__core-js_shared__",
		"_countAA",
		"ae",
		"alert",
		"api_target",
		"api_target_sri",
		"ark",
		"async_fingerprints",
		"atob",
		"blur",
		"btoa",
		"caches",
		"cancelAnimationFrame",
		"cancelIdleCallback",
		"capiMode",
		"capiSettings",
		"capiVersion",
		"captureEvents",
		"cdn",
		"chrome",
		"clearInterval",
		"clearTimeout",
		"clientInformation",
		"close",
		"closed",
		"confirm",
		"console",
		"cookieStore",
		"createImageBitmap",
		"credentialless",
		"crossOriginIsolated",
		"crypto",
		"customElements",
		"decodeURI",
		"decodeURIComponent",
		"devicePixelRatio",
		"doBBBd",
		"document",
		"encodeURI",
		"encodeURIComponent",
		"escape",
		"eval",
		"event",
		"extended_fingerprinting_enabled",
		"external",
		"fc_api_server",
		"fc_fp",
		"fc_obj",
		"fetch",
		"find",
		"find_onload",
		"fingerprinting_enabled",
		"focus",
		"fp_result",
		"frameElement",
		"frames",
		"getComputedStyle",
		"getScreenDetails",
		"getSelection",
		"get_outer_html",
		"get_query_data",
		"globalThis",
		"history",
		"indexedDB",
		"innerHeight",
		"innerWidth",
		"isFinite",
		"isNaN",
		"isSecureContext",
		"launchQueue",
		"length",
		"loadedWithData",
		"localStorage",
		"location",
		"locationbar",
		"log",
		"matchMedia",
		"menubar",
		"moveBy",
		"moveTo",
		"msie",
		"name",
		"navigation",
		"navigator",
		"offscreenBuffering",
		"onabort",
		"onafterprint",
		"onanimationend",
		"onanimationiteration",
		"onanimationstart",
		"onappinstalled",
		"onauxclick",
		"onbeforeinput",
		"onbeforeinstallprompt",
		"onbeforematch",
		"onbeforeprint",
		"onbeforetoggle",
		"onbeforeunload",
		"onbeforexrselect",
		"onblur",
		"oncancel",
		"oncanplay",
		"oncanplaythrough",
		"onchange",
		"onclick",
		"onclose",
		"oncontentvisibilityautostatechange",
		"oncontextlost",
		"oncontextmenu",
		"oncontextrestored",
		"oncuechange",
		"ondblclick",
		"ondevicemotion",
		"ondeviceorientation",
		"ondeviceorientationabsolute",
		"ondrag",
		"ondragend",
		"ondragenter",
		"ondragleave",
		"ondragover",
		"ondragstart",
		"ondrop",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onfocus",
		"onformdata",
		"ongotpointercapture",
		"onhashchange",
		"oninput",
		"oninvalid",
		"onkeydown",
		"onkeypress",
		"onkeyup",
		"onlanguagechange",
		"onload",
		"onload_retry",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onlostpointercapture",
		"onmessage",
		"onmessageerror",
		"onmousedown",
		"onmouseenter",
		"onmouseleave",
		"onmousemove",
		"onmouseout",
		"onmouseover",
		"onmouseup",
		"onmousewheel",
		"onoffline",
		"ononline",
		"onpagehide",
		"onpageshow",
		"onpause",
		"onplay",
		"onplaying",
		"onpointercancel",
		"onpointerdown",
		"onpointerenter",
		"onpointerleave",
		"onpointermove",
		"onpointerout",
		"onpointerover",
		"onpointerrawupdate",
		"onpointerup",
		"onpopstate",
		"onprogress",
		"onratechange",
		"onrejectionhandled",
		"onreset",
		"onresize",
		"onscroll",
		"onscrollend",
		"onsearch",
		"onsecuritypolicyviolation",
		"onseeked",
		"onseeking",
		"onselect",
		"onselectionchange",
		"onselectstart",
		"onslotchange",
		"onstalled",
		"onstorage",
		"onsubmit",
		"onsuspend",
		"ontimeupdate",
		"ontoggle",
		"ontransitioncancel",
		"ontransitionend",
		"ontransitionrun",
		"ontransitionstart",
		"onunhandledrejection",
		"onunload",
		"onvolumechange",
		"onwaiting",
		"onwebkitanimationend",
		"onwebkitanimationiteration",
		"onwebkitanimationstart",
		"onwebkittransitionend",
		"onwheel",
		"open",
		"openDatabase",
		"opener",
		"origin",
		"originAgentCluster",
		"outerHeight",
		"outerWidth",
		"pageXOffset",
		"pageYOffset",
		"parent",
		"parseFloat",
		"parseInt",
		"performance",
		"personalbar",
		"postMessage",
		"print",
		"prompt",
		"public_key",
		"queryLocalFonts",
		"query_data",
		"queueMicrotask",
		"releaseEvents",
		"reportError",
		"requestAnimationFrame",
		"requestIdleCallback",
		"resizeBy",
		"resizeTo",
		"scheduler",
		"screen",
		"screenLeft",
		"screenTop",
		"screenX",
		"screenY",
		"scroll",
		"scrollBy",
		"scrollTo",
		"scrollX",
		"scrollY",
		"scrollbars",
		"self",
		"sessionStorage",
		"setAPIInput",
		"setInterval",
		"setQueryDataInput",
		"setTimeout",
		"showDirectoryPicker",
		"showOpenFilePicker",
		"showSaveFilePicker",
		"siteData",
		"speechSynthesis",
		"startArkoseEnforcement",
		"status",
		"statusbar",
		"stop",
		"stringifyWithFloat",
		"structuredClone",
		"styleMedia",
		"target",
		"toolbar",
		"top",
		"trustedTypes",
		"undefined",
		"unescape",
		"visualViewport",
		"webkitCancelAnimationFrame",
		"webkitMediaStream",
		"webkitRTCPeerConnection",
		"webkitRequestAnimationFrame",
		"webkitRequestFileSystem",
		"webkitResolveLocalFileSystemURL",
		"webkitSpeechGrammar",
		"webkitSpeechGrammarList",
		"webkitSpeechRecognition",
		"webkitSpeechRecognitionError",
		"webkitSpeechRecognitionEvent",
		"webkitURL",
		"window",
	}
	sort.Strings(b1)
	result := strings.Join(b1, "|")
	return getMurmur128String(result, 420)
}

func getWindowProtoChainHash() string { // return this[dh(f_a_hU.f)](b0[dh(f_a_hU.g)]('|'), 0x1a4);
	// Object.getPrototypeOf(window);
	b0 := []string{
		"TEMPORARY",
		"PERSISTENT",
		"constructor",
		"addEventListener",
		"dispatchEvent",
		"removeEventListener",
		"constructor",
		"constructor",
		"__defineGetter__",
		"__defineSetter__",
		"hasOwnProperty",
		"__lookupGetter__",
		"__lookupSetter__",
		"isPrototypeOf",
		"propertyIsEnumerable",
		"toString",
		"valueOf",
		"__proto__",
		"toLocaleString",
	}
	result2 := strings.Join(b0, "|")
	return x64hash128(result2, 420)
}

func x64Add(m []uint32, n []uint32) []uint32 {
	m = []uint32{m[0] >> 16, m[0] & 0xffff, m[1] >> 16, m[1] & 0xffff}
	n = []uint32{n[0] >> 16, n[0] & 0xffff, n[1] >> 16, n[1] & 0xffff}
	o := []uint32{0, 0, 0, 0}
	o[3] += m[3] + n[3]
	o[2] += o[3] >> 16
	o[3] &= 0xffff
	o[2] += m[2] + n[2]
	o[1] += o[2] >> 16
	o[2] &= 0xffff
	o[1] += m[1] + n[1]
	o[0] += o[1] >> 16
	o[1] &= 0xffff
	o[0] += m[0] + n[0]
	o[0] &= 0xffff
	return []uint32{(o[0] << 16) | o[1], (o[2] << 16) | o[3]}
}

func x64Multiply(m []uint32, n []uint32) []uint32 {
	m = []uint32{m[0] >> 16, m[0] & 0xffff, m[1] >> 16, m[1] & 0xffff}
	n = []uint32{n[0] >> 16, n[0] & 0xffff, n[1] >> 16, n[1] & 0xffff}
	o := []uint32{0, 0, 0, 0}
	o[3] += m[3] * n[3]
	o[2] += o[3] >> 16
	o[3] &= 0xffff
	o[2] += m[2] * n[3]
	o[1] += o[2] >> 16
	o[2] &= 0xffff
	o[2] += m[3] * n[2]
	o[1] += o[2] >> 16
	o[2] &= 0xffff
	o[1] += m[1] * n[3]
	o[0] += o[1] >> 16
	o[1] &= 0xffff
	o[1] += m[2] * n[2]
	o[0] += o[1] >> 16
	o[1] &= 0xffff
	o[1] += m[3] * n[1]
	o[0] += o[1] >> 16
	o[1] &= 0xffff
	o[0] += m[0]*n[3] + m[1]*n[2] + m[2]*n[1] + m[3]*n[0]
	o[0] &= 0xffff
	return []uint32{(o[0] << 16) | o[1], (o[2] << 16) | o[3]}
}

//goland:noinspection SpellCheckingInspection
func x64Rotl(m []uint32, n uint32) []uint32 {
	n %= 64
	if n == 32 {
		return []uint32{m[1], m[0]}
	} else if n < 32 {
		return []uint32{(m[0] << n) | (m[1] >> (32 - n)), (m[1] << n) | (m[0] >> (32 - n))}
	} else {
		n -= 32
		return []uint32{(m[1] << n) | (m[0] >> (32 - n)), (m[0] << n) | (m[1] >> (32 - n))}
	}
}

func x64LeftShift(m []uint32, n uint32) []uint32 {
	n %= 64
	if n == 0 {
		return m
	} else if n < 32 {
		return []uint32{(m[0] << n) | (m[1] >> (32 - n)), m[1] << n}
	} else {
		return []uint32{m[1] << (n - 32), 0}
	}
}

func x64Xor(m []uint32, n []uint32) []uint32 {
	return []uint32{m[0] ^ n[0], m[1] ^ n[1]}
}

//goland:noinspection SpellCheckingInspection
func x64Fmix(h []uint32) []uint32 {
	h = x64Xor(h, []uint32{0, h[0] >> 1})
	h = x64Multiply(h, []uint32{0xff51afd7, 0xed558ccd})
	h = x64Xor(h, []uint32{0, h[0] >> 1})
	h = x64Multiply(h, []uint32{0xc4ceb9fe, 0x1a85ec53})
	h = x64Xor(h, []uint32{0, h[0] >> 1})
	return h
}

func x64hash128(key string, seed uint32) string {
	keyLength := len(key)
	remainder := keyLength % 16
	bytes := keyLength - remainder

	var h1 = []uint32{0, seed}
	var h2 = []uint32{0, seed}
	var k1 = []uint32{0, 0}
	var k2 = []uint32{0, 0}
	var c1 = []uint32{0x87c37b91, 0x114253d5}
	var c2 = []uint32{0x4cf5ad43, 0x2745937f}

	for i := 0; i < bytes; i += 16 {
		k1[0] = uint32(key[i+4])&0xff | (uint32(key[i+5])&0xff)<<8 | (uint32(key[i+6])&0xff)<<16 | (uint32(key[i+7])&0xff)<<24
		k1[1] = uint32(key[i])&0xff | (uint32(key[i+1])&0xff)<<8 | (uint32(key[i+2])&0xff)<<16 | (uint32(key[i+3])&0xff)<<24

		k2[0] = uint32(key[i+12])&0xff | (uint32(key[i+13])&0xff)<<8 | (uint32(key[i+14])&0xff)<<16 | (uint32(key[i+15])&0xff)<<24
		k2[1] = uint32(key[i+8])&0xff | (uint32(key[i+9])&0xff)<<8 | (uint32(key[i+10])&0xff)<<16 | (uint32(key[i+11])&0xff)<<24

		k1 = x64Multiply(k1, c1)
		k1 = x64Rotl(k1, 31)
		k1 = x64Multiply(k1, c2)
		h1 = x64Xor(h1, k1)
		h1 = x64Rotl(h1, 27)
		h1 = x64Add(h1, h2)
		h1 = x64Add(x64Multiply(h1, []uint32{0, 5}), []uint32{0, 0x52dce729})

		k2 = x64Multiply(k2, c2)
		k2 = x64Rotl(k2, 33)
		k2 = x64Multiply(k2, c1)
		h2 = x64Xor(h2, k2)
		h2 = x64Rotl(h2, 31)
		h2 = x64Add(h2, h1)
		h2 = x64Add(x64Multiply(h2, []uint32{0, 5}), []uint32{0, 0x38495ab5})
	}

	k1 = []uint32{0, 0}
	k2 = []uint32{0, 0}

	switch remainder {
	case 15:
		k2 = x64Xor(k2, x64LeftShift([]uint32{0, uint32(key[bytes+14])}, 48))
		fallthrough
	case 14:
		k2 = x64Xor(k2, x64LeftShift([]uint32{0, uint32(key[bytes+13])}, 40))
		fallthrough
	case 13:
		k2 = x64Xor(k2, x64LeftShift([]uint32{0, uint32(key[bytes+12])}, 32))
		fallthrough
	case 12:
		k2 = x64Xor(k2, x64LeftShift([]uint32{0, uint32(key[bytes+11])}, 24))
		fallthrough
	case 11:
		k2 = x64Xor(k2, x64LeftShift([]uint32{0, uint32(key[bytes+10])}, 16))
		fallthrough
	case 10:
		k2 = x64Xor(k2, x64LeftShift([]uint32{0, uint32(key[bytes+9])}, 8))
		fallthrough
	case 9:
		k2 = x64Xor(k2, []uint32{0, uint32(key[bytes+8])})
		k2 = x64Multiply(k2, c2)
		k2 = x64Rotl(k2, 33)
		k2 = x64Multiply(k2, c1)
		h2 = x64Xor(h2, k2)
		fallthrough
	case 8:
		k1 = x64Xor(k1, x64LeftShift([]uint32{0, uint32(key[bytes+7])}, 56))
		fallthrough
	case 7:
		k1 = x64Xor(k1, x64LeftShift([]uint32{0, uint32(key[bytes+6])}, 48))
		fallthrough
	case 6:
		k1 = x64Xor(k1, x64LeftShift([]uint32{0, uint32(key[bytes+5])}, 40))
		fallthrough
	case 5:
		k1 = x64Xor(k1, x64LeftShift([]uint32{0, uint32(key[bytes+4])}, 32))
		fallthrough
	case 4:
		k1 = x64Xor(k1, x64LeftShift([]uint32{0, uint32(key[bytes+3])}, 24))
		fallthrough
	case 3:
		k1 = x64Xor(k1, x64LeftShift([]uint32{0, uint32(key[bytes+2])}, 16))
		fallthrough
	case 2:
		k1 = x64Xor(k1, x64LeftShift([]uint32{0, uint32(key[bytes+1])}, 8))
		fallthrough
	case 1:
		k1 = x64Xor(k1, []uint32{0, uint32(key[bytes])})
		k1 = x64Multiply(k1, c1)
		k1 = x64Rotl(k1, 31)
		k1 = x64Multiply(k1, c2)
		h1 = x64Xor(h1, k1)
	}

	h1 = x64Xor(h1, []uint32{0, uint32(keyLength)})
	h2 = x64Xor(h2, []uint32{0, uint32(keyLength)})
	h1 = x64Add(h1, h2)
	h2 = x64Add(h2, h1)
	h1 = x64Fmix(h1)
	h2 = x64Fmix(h2)
	h1 = x64Add(h1, h2)
	h2 = x64Add(h2, h1)

	return fmt.Sprintf("%08x%08x%08x%08x", h1[0], h1[1], h2[0], h2[1])
}

func getCFPHash(cfp string) uint32 {
	//'this is the cfp: canvas xxx base64 image'.split('').reduce((b5, b6) => {
	//	return b5 = (b5 << 5) - b5 + b6.charCodeAt(0), b5 & b5;
	//}, 0);

	var b5 uint32
	for _, b6 := range cfp {
		b5 = (b5 << 5) - b5 + uint32(b6)
		b5 &= b5
	}
	return b5
}

func getIfeHash() string {
	return x64hash128(strings.Join(getFeList(), ", "), 38)
}