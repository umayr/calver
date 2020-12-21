const {Binary} = require('binary-install');
const os = require('os');

// Mapping between Node's `process.platform` to Golang's
const PLATFORM_MAPPING = {
    "darwin": "darwin",
    "linux": "linux",
    "win32": "windows",
};

// Mapping from Node's `process.arch` to Golang's `$GOARCH`
const ARCH_MAPPING = {
    "ia32": "386",
    "x64": "amd64",
    "arm": "arm"
};

function getPlatform() {
    const type = os.type().toLowerCase();
    const arch = os.arch().toLowerCase();

    if (typeof PLATFORM_MAPPING[type] === 'undefined' || typeof ARCH_MAPPING[arch] === 'undefined') {
        throw new Error(`Unsupported platform: ${type} ${arch}`);
    }

    return {os: PLATFORM_MAPPING[type], arch: ARCH_MAPPING[arch]};
}

function getBinary() {
    const {os, arch} = getPlatform();
    const version = require('../package.json').version;

    const name = 'calver';
    const url = `https://github.com/umayr/${name}/releases/download/${version}/${name}-${os}-${arch}.tar.gz`;

    return new Binary(name, url);
}

module.exports = getBinary;