<template>
<v-card :dark="useDarkTheme">
  <v-container>
    <v-layout row wrap>
      <v-flex xs1>
        <v-checkbox v-model="completed" @click.stop="isCompletedHandler();"></v-checkbox>
      </v-flex>

      <v-flex xs10>
        <v-card-title>
          <div>
            <div class="headline">{{ name }}</div>
            <em>Due on {{ date }} at {{ time }}</em>
          </div>
        </v-card-title>

        <v-card-text>
          <v-layout row wrap>
            <v-flex xs12 v-if="description !== ''">
              <p>{{ description }}</p>
            </v-flex>
          </v-layout>
        </v-card-text>
      </v-flex>

      <v-flex xs1>
        <v-layout fill-height column justify-space-between align-end>
          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-icon class="mr-2" v-on="on" @click.stop="remove();">
                far fa-times
              </v-icon>
            </template>
            <span>Delete</span>
          </v-tooltip>

          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-icon class="mb-2" v-on="on" @click.stop="edit();">
                far fa-edit
              </v-icon>
            </template>
            <span>Edit</span>
          </v-tooltip>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</v-card>
</template>

<script>
import { mapActions, mapMutations } from 'vuex';
import axios from 'axios';

import taskMixin from '@/mixins/task';

export default {
  name: 'task',
  mixins: [taskMixin],
  computed: {
    useDarkTheme() {
      return this.theme === 'dark';
    },
  },
  data() {
    return {
      completed: this.category === 'completed',
    };
  },
  methods: {
    ...mapMutations('status', [
      'setMainError',
    ]),
    ...mapMutations('task', [
      'deleteTask',
    ]),
    ...mapMutations('taskForm', [
      'openTaskForm',
    ]),
    ...mapActions('task', [
      'updateTask',
    ]),
    async remove() {
      const resp = await fetch('/', response => response);
      axios.post('/delete_task', {
        index: this.index,
        name: this.name,
        date: this.date,
        time: this.time,
        description: this.description,
        category: this.category,
      }, {
        headers: { 'X-CSRF-TOKEN': resp.headers.get('X-CSRF-TOKEN') },
      }).then((response) => {
        if (response.data.statusType === 'success') {
          this.deleteTask({
            category: this.category,
            position: this.position,
          });
        } else {
          this.setMainError(response.data.status);
        }
      }).catch((error) => {
        this.setMainError(error);
      });
    },
    edit() {
      this.openTaskForm({
        index: this.index,
        name: this.name,
        date: this.date,
        time: this.time,
        description: this.description,
        category: this.category,

        position: this.position,
        action: 'edit',
      });
    },
    async isCompletedHandler() {
      const task = {
        index: this.index,
        name: this.name,
        date: this.date,
        time: this.time,
        description: this.description,
        category: this.category,
        position: this.position,

        oldCategory: this.category,
        oldPosition: this.position,
      };
      if (!this.completed) {
        task.category = 'completed';
        this.updateTask(task);
      } else {
        this.assignCategory(task);
        this.updateTask(task);
      }

      const resp = await fetch('/', response => response);
      axios.post(`/is_completed/${!this.completed}`, task, {
        headers: { 'X-CSRF-TOKEN': resp.headers.get('X-CSRF-TOKEN') },
      }).then((response) => {
        if (response.data.statusType !== 'success') {
          this.setMainError(response.data.status);
        }
      }).catch((error) => {
        this.setMainError(error);
      });
    },
  },
  props: {
    index: Number,
    name: String,
    date: String,
    time: String,
    description: String,
    category: String,
    position: Number,

    theme: {
      default: 'light',
      type: String,
      validator: value => ['light', 'dark'].indexOf(value) !== -1,
    },
  },
};
</script>
