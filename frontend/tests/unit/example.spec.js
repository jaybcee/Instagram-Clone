import { shallowMount } from '@vue/test-utils'
import Comment from '@/components/Comment.vue'

describe('test', () => {
  it('test', () => {
    const bool = true;
    const wrapper = shallowMount(Comment, {});
    expect(bool).toMatch(true)
  })
})
