function lines2Int(data) {
    var ints = []
    lines = data.trim().split('\n');
    for (l in lines){
        ints.push(parseInt(lines[l], 10));
    }
    return ints;
}

function increases(ints) {
    sum = 0;
    for (let i=0; i < ints.length; i++) {
        if (i != 0) {
            if (ints[i] > ints[i-1]){
                sum += 1;
            } 
        }
    }
    return sum;
}

function threes(ints) {
    var t = []
    for (let i=0; i < ints.length -2; i++){ 
            console.log(i, ints[i], ints[i+1], ints[i+2]);
            t.push(ints[i] + ints[i+1] + ints[i+2]);
    }

    return t;
}



async function run() {
    var sample = `
    199
    200
    208
    210
    200
    207
    240
    269
    260
    263
    `

    var ints = await fetch("https://adventofcode.com/2021/day/1/input", {
      "headers": {
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
        "accept-language": "en-US,en;q=0.9",
        "sec-ch-ua": "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"",
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": "\"Windows\"",
        "sec-fetch-dest": "document",
        "sec-fetch-mode": "navigate",
        "sec-fetch-site": "same-origin",
        "sec-fetch-user": "?1",
        "upgrade-insecure-requests": "1"
      },
      "referrer": "https://adventofcode.com/2021/day/1",
      "referrerPolicy": "strict-origin-when-cross-origin",
      "body": null,
      "method": "GET",
      "mode": "cors",
      "credentials": "include"
    })
    .then(res => res.text())
    .then(data => lines2Int(data));

    //var aoc1a = increases(ints);
    console.log(ints);
    var aoc1b = increases(threes(ints));

    //console.log(aoc1a);
    console.log(aoc1b);
}

run();
