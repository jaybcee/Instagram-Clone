<template>
  <v-card max-width="450" class="mx-auto">
    <v-list three-line>
      <v-subheader v-text="'Comments'"></v-subheader>
      <template v-for="(item, index) in items">
        <v-divider :key="index" :inset="item.inset"></v-divider>

        <v-list-item :key="item.uuid">
          <v-list-item-avatar>
            <v-img :src="item.avatar"></v-img>
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title v-text="item.username"></v-list-item-title>
            <v-list-item-subtitle v-text="item.subusername"></v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
    <v-container>
      <v-row>
        <v-col cols="8">
          <v-text-field label="Enter a comment" single-line filled v-model="comment"></v-text-field>
        </v-col>
        <v-col cols="4">
          <v-btn @click="addComment" small color="primary">Primary</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import uuid from 'uuid';
export default {
  data: () => ({
    comment: "",
    items: [
      {
        avatar: "https://cdn.vuetifyjs.com/images/lists/1.jpg",
        username: "Tim",
        subusername: "whatever",
        time: 1
      },

      {
        avatar: "https://cdn.vuetifyjs.com/images/lists/2.jpg",
        username: "Tom",
        subusername: "whatever",
        time: 3
      },

      {
        avatar: "https://cdn.vuetifyjs.com/images/lists/3.jpg",
        username: "Rolling stones",
        subusername: "whatever",
        time: 0
      },

      {
        avatar: "https://cdn.vuetifyjs.com/images/lists/4.jpg",
        username: "Tam",
        subusername: "whatever",
        time: 900
      },

      {
        avatar: "https://cdn.vuetifyjs.com/images/lists/5.jpg",
        username: "Tum",
        subusername: "whatever",
        time: 314
      }
    ]
  }),
  methods: {
    addComment() {
      console.log(this.comment);
      let c = {
        avatar: "https://cdn.vuetifyjs.com/images/lists/5.jpg",
        username: "Paul",
        subusername: this.comment,
        time: -1,
        uuid: uuid.v4()
      };
      this.items.push(c);
    }
  },
  mounted() {
    console.log(this.items);
    for (let i = 0;i<i.length;i++) {
        this.items[i].uuid=uuid.v4()
    }
    this.items.sort(function(a, b) {
      return a.time - b.time;
    });
    
  }
};
</script>