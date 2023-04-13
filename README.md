# cgpt_scaleWeightsAndMeasures_Go
OpenAI API Weights and Measures REST service in Go

#### Weights and Measures scaling using OpenAI API - Demo

Caution - this is just a demo - there is not much error checking !

This is a simple/toy REST service in Go using OpenAI API for converting and scaling weights and measures (it can be used for other purposes with a few changes).

You will need Go (please see the references).

Get OpenAI API key (check References towards the end). Copy `.env.example` to `.env` and add the API key there.

After installing `Go` , it should be possible to run the demo with something like `go run main.go`. This will start a local REST service on port 3000.

To try the backend REST service - it should be possible to test some weights and measures scaling REST endpoint with something like:

`http://localhost:3000/scaleWeightsAndMeasures?quantity=1&fromUnit=kilogram&toUnit=pound&scaleFactor=2`

from a browser. 

OR

`curl 'http://localhost:3000/scaleWeightsAndMeasures?quantity=1&fromUnit=kilogram&toUnit=pound&scaleFactor=2'`

from commandline.

It will comeup with an answer like:

```
Converted quantity : {
    "choices": [
        {
            "finish_reason": "stop",
            "index": 0,
            "logprobs": null,
            "text": "\n2.00 kg = 4.40924524 pounds"
        }
    ],
    "created": 1681361203,
    "id": "cmpl-74j91ZpnR6mkgquW7jOeIOng8sEHE",
    "model": "text-davinci-003",
    "object": "text_completion",
    "usage": {
        "completion_tokens": 12,
        "prompt_tokens": 209,
        "total_tokens": 221
    }
}
```

Experiment with training data and question !

---
#### References
1. [OpenAI quickstart](https://platform.openai.com/docs/quickstart)
2. [Imperial units](https://en.wikipedia.org/wiki/Imperial_units)
3. [Go everything](https://go.dev/)
