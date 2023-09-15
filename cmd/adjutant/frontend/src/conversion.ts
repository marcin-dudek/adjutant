
/// converts bytes to MB
export const toMB = (n: number) => (n / 1_000_000).toFixed(2);

const pad = (d: number) => d.toString().padStart(2, "0");

/// converts number of ns to duration in format "[xx]h [yy]min [zz]s"
export const toDuration = (duration: number) => {
    let d = duration / 1_000_000_000;
    let seconds: number = d % 60;
    let minutes: number = ((d - seconds) / 60) % 60;
    let hours: number = (d - minutes * 60 - seconds) / (60 * 60);

    if (hours > 0) {
        return `${hours}h ${pad(minutes)}min ${pad(seconds)}s`;
    } else if (minutes > 0) {
        return `${minutes}min ${pad(seconds)}s`;
    } else {
        return `${seconds}s`;
    }
};