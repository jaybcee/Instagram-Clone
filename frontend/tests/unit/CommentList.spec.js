import { shallowMount } from '@vue/test-utils'
import CommentList from '@/components/CommentList.vue'

describe('CommentList', () => {
  test('is a Vue instance', () => {
    const wrapper = shallowMount(CommentList)
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})