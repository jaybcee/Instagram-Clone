module.exports = {
    extends: [
        // add more generic rulesets here, such as:
        'eslint:recommended',
        'plugin:vue/recommended'
    ],
    rules: {
        // override/add rules settings here, such as:
        // "linebreak-style": ["warn", "windows"]
        "vue/html-self-closing": 0,
    },
    "globals": {
        "process": true
    }
}