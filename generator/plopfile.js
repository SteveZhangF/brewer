import modelGenerator from './plop-templates/model/prompt.js'

export default function (plop) {
  plop.setGenerator('model', modelGenerator)
}