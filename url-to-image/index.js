const { chromium } = require('playwright-core')
const express = require('express')
const app = express()

app.use(express.json())
app.post('/*', async (req, res) => {
    browser = await chromium.launch()
    page = await browser.newPage()
    await page.goto(req.body.url)
    image = await page.screenshot({
        fullPage: true
    })

    res.json({
        content: [{
            image: {
                contentType: "image/png",
                base64: image.toString('base64')
            }
        }]
    })

    await browser.close()
});

app.listen(8080)