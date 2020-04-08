import { mount } from '@vue/test-utils'
// import Comment from '../components/Comment'
// import Comment from "./Comment";
import Comment from "../components/Comment.vue"

describe('Comment', () => {
  test('is a Vue instance', () => {
    const wrapper = mount(Comment)
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})