const express = require('express')
const path = require('path')
const axios = require('axios')
const app = express()

app.set('views', path.join(__dirname, 'views'))
app.set('view engine', 'ejs')

/*
app.get('/', (req, res) => {
    res.render('index')
})
*/

app.get('/', (req, res) => {
    res.send("Hello World")
})

app.get('/getGoogle', (req, res) => {
    axios
        .get('http://www.google.com')
        .then(output => {
            res.send(output.status)
        })
        .catch(error => {
            console.log(error)
        })
})

app.get('/getGithub', (req, res) => {
    axios
        .get('http://www.github.com')
        .then(output => {
            res.send(output.status)
        })
        .catch(error => {
            console.log(error)
        })
})

app.listen(3000, () => {
    console.log('Server listening to port 3000')
})