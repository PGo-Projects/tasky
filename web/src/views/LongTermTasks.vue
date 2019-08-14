<template>
<v-container fluid fill-height>
  <v-layout row align-start justify-center v-if="longTermTasks.length !== 0">
    <v-flex xs12>
      <div v-for="(task, index) in longTermTasks" :key="task.id">
        <task :index="task.index" :name="task.name" :date="task.date" :time="task.time"
              :description="task.description" :category="task.category" :position="index">
        </task>
      </div>
    </v-flex>
  </v-layout>
  <v-layout row align-center justify-center v-else>
    <v-flex xs12>
      <p class="display-3 text-xs-center">You don't have any long term tasks!</p>
    </v-flex>
  </v-layout>
</v-container>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex';
import axios from 'axios';

import Task from '@/components/Task.vue';

export default {
  components: { Task },
  name: 'longtermtasks',
  mounted() {
    axios.get('/get_category/longTerm').then((resp) => {
      const tasks = JSON.parse(resp.data.tasks);
      this.setTasks({
        category: 'longTerm',
        tasks,
      });
    });
  },
  computed: {
    ...mapGetters('task', [
      'longTermTasks',
    ]),
  },
  methods: {
    ...mapMutations('task', [
      'setTasks',
    ]),
  },
};
</script>
