const fs = require('fs');

function castInt(s) {
    const num = parseInt(s, 10);
    if (isNaN(num)) {
        throw new Error('Error casting string to int');
    }
    return num;
}

function removeAt(array, index) {
    if (array.length === 0) {
        console.error('Error: cannot remove from an empty array');
        return array;
    }
    if (index < 0 || index >= array.length) {
        console.error(`Error: index ${index} out of range (array length ${array.length})`);
        return array;
    }
    return array.slice(0, index).concat(array.slice(index + 1));
}

function isSafe(i, levels, pos, increasing, canTolerate) {
    const current = levels[pos];
    const next = levels[pos + 1];

    if ((next > current && next - current <= 3) && (increasing + 1 <= 2)) {
        if (pos === levels.length - 2) {
            return true;
        }
        return isSafe(i, levels, pos + 1, 1, canTolerate);
    } else if ((current > next && current - next <= 3) && (increasing % 2 === 0)) {
        if (pos === levels.length - 2) {
            return true;
        }
        return isSafe(i, levels, pos + 1, 2, canTolerate);
    } else if (canTolerate) {
        const safeWOL1 = isSafe(i, removeAt(levels, pos), 0, 0, false);
        const safeWOL2 = isSafe(i, removeAt(levels, pos + 1), 0, 0, false);
        return safeWOL1 || safeWOL2;
    }
    return false;
}

function solutionPart1() {
    const file = fs.readFileSync('input.txt', 'utf-8');
    const lines = file.split('\n').filter(line => line.trim().length > 0);

    let safe = 0;

    for (const line of lines) {
        const strLevels = line.split(' ');

        const l1 = castInt(strLevels[0]);
        const l2 = castInt(strLevels[1]);

        let increasing = false;
        if (l2 > l1 && l2 - l1 <= 3) {
            increasing = true;
        } else if (l1 > l2 && l1 - l2 <= 3) {
            increasing = false;
        } else {
            continue;
        }

        let prevLevel = l2;
        let isSafe = true;

        for (const s of strLevels.slice(2)) {
            const level = castInt(s);
            const diff = Math.abs(level - prevLevel);
            if ((level > prevLevel && diff <= 3 && increasing) || (prevLevel > level && diff <= 3 && !increasing)) {
                prevLevel = level;
            } else {
                isSafe = false;
                break;
            }
        }

        if (isSafe) {
            safe++;
        }
    }

    return safe;
}

function solutionPart2() {
    const file = fs.readFileSync('input.txt', 'utf-8');
    const lines = file.split('\n').filter(line => line.trim().length > 0);

    let safe = 0;
    let i = 1;

    for (const line of lines) {
        const strLevels = line.split(' ');
        const levels = strLevels.map(castInt);

        const levelsCopy = [...levels];

        if (isSafe(i, levelsCopy, 0, 0, true)) {
            safe++;
        } else {
            const l3 = levels.slice(0, levels.length - 1);
            const s1 = isSafe(i, removeAt(levels, 0), 0, 0, false);
            const s2 = isSafe(i, removeAt(levels, 1), 0, 0, false);
            const s3 = isSafe(i, l3, 0, 0, false);

            if (s1 || s2 || s3) {
                safe++;
            }
        }
        i++;
    }

    return safe;
}

function main() {
    console.log('sol1:', solutionPart1());
    console.log('sol2:', solutionPart2());
}

main();
