const generate = async () => {
    const explanation = document.getElementById("feature-explanation").value;
    const res = await fetch("/generate/scenarios/streaming", {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            "question": explanation
        })
    })
    // console.log(res)
    const reader = res.body.getReader()
    // console.log(reader)

    while (true) {
        const { done, value } = await reader.read()
        // console.log(done)
        //console.log(value)
        let chunk = String.fromCharCode.apply(null, value);
        // chunk = chunk.replaceAll("data: ", "")
        // chunk.
        // let chunks = chunk.split("\n\n")
        // chunks.map(chunk => chunk = chunk.replaceAll("\n", ""))
        console.log("------------------")
        console.log(chunk)
        console.log("------------------")

        

        // chunk = chunk.trimEnd()
        // chunk = chunk.replace("[", "")
        // chunk = chunk.replace("]", "")
        // // console.log(chunk)
        document.getElementById("scenarios").innerHTML += chunk
        if (done) break;
    }
}