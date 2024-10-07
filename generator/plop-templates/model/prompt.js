import { notEmpty } from '../utils.js'
export default {
  description: 'generate new model',
  prompts: [{
    type: 'input',
    name: 'name',
    message: 'model name please',
    validate: notEmpty('name')
  },
  ],
  actions: data => {
    const name = '{{dashCase name}}'
    const propName = '{{properCase name}}'
    const actions = [
      {
        name: 'index',
        type: 'add',
        path: `../model/{{snakeCase name}}.go`,
        templateFile: 'plop-templates/model/model.hbs',
        data: {
          name: name
        }
      },
      {
        name: 'router',
        type: 'modify',
        path: '../main.go',
        pattern: /\/\/ ADD MORE ROUTERS HERE DO NOT DELETE THIS LINE/gi,
        templateFile: 'plop-templates/model/router.hbs'
      },
      {
        name: 'store',
        type: 'add',
        path: `../storage/{{snakeCase name}}.go`,
        templateFile: 'plop-templates/model/store.hbs',
        data: {
          name: name
        }
      },
      {
        name: 'command',
        type: 'add',
        path: `../command/{{snakeCase name}}.go`,
        templateFile: 'plop-templates/model/command.hbs',
        data: {
          name: name
        }
      }
    ].filter(action => {
      return true
    })

    return actions
  }
}
