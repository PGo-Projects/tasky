<template>
<v-dialog persistent v-model="openTaskForm" max-width="1000px">
  <v-card>
    <v-container>
      <v-card-title>
        <span class="headline" :color="headerColor">{{ headerText }}</span>
      </v-card-title>

      <div class="my-4">
        <v-alert v-if="status !== undefined" :type="statusType"
                 value=true outline=true>
          {{ status }}
        </v-alert>
      </div>

      <v-card-text>
        <v-form ref="form" @submit.prevent="saveForm">
          <v-layout row wrap>
            <v-flex xs12>
              <v-text-field
                v-model="info.name"
                :rules="nameRules"
                :value="info.name"
                :color="formColor"
                label="Name"
                name="name"
                ref="name"
                required
                ></v-text-field>
            </v-flex>

            <v-flex xs12 md5>
              <v-menu
                ref="dateMenu"
                v-model="dateMenu"
                :color="formColor"
                :close-on-content-click="false"
                :return-value.sync="info.date"
                lazy
                transition="scale-transition"
                offset-y
                full-width
                min-width="280px"
                >
                <template v-slot:activator="{ on }">
                  <v-text-field
                    v-model="info.date"
                    :rules="dateRules"
                    :value="info.date"
                    :color="formColor"
                    label="Date"
                    prepend-icon="far fa-calendar-day"
                    readonly
                    @click="activateDatepicker"
                    @focus="dateMenu = true;"
                    ></v-text-field>
                </template>
                <v-date-picker v-model="info.date" :color="formColor"
                               :dark="useDarkTheme" no-title scrollable>
                  <v-spacer></v-spacer>
                  <v-btn flat :color="formColor" @click="dateMenu = false;">
                    Cancel
                  </v-btn>
                  <v-btn flat :color="formColor"
                         @click="$refs.dateMenu.save(info.date)">
                    OK
                  </v-btn>
                </v-date-picker>
              </v-menu>
            </v-flex>
            <v-flex md1>
            </v-flex>
            <v-flex xs12 md5>
              <v-menu
                ref="timeMenu"
                v-model="timeMenu"
                :color="formColor"
                :close-on-content-click="false"
                :return-value.sync="info.time"
                lazy
                transition="scale-transition"
                offset-y
                full-width
                min-width="280px"
                @input="value"
                >
                <template v-slot:activator="{ on }">
                  <v-text-field
                    v-model="info.time"
                    :rules="timeRules"
                    :value="info.time"
                    :color="formColor"
                    label="Time"
                    prepend-icon="far fa-clock"
                    readonly
                    @click="activateTimepicker"
                    @focus="timeMenu = true;"
                    ></v-text-field>
                </template>
                <v-time-picker v-if="taskOpenForm" v-model="info.time" ref="timePicker"
                               :color="formColor" :dark="useDarkTheme"
                               scrollable ampm-in-title>
                  <v-spacer></v-spacer>
                  <v-btn flat :color="formColor" @click="timeMenu = false;">
                    Cancel
                  </v-btn>
                  <v-btn flat :color="formColor"
                         @click="$refs.timeMenu.save(info.time)">
                    OK
                  </v-btn>
                </v-time-picker>
              </v-menu>
            </v-flex>

            <v-flex xs12>
              <v-textarea
                v-model="info.description"
                :value="info.description"
                :color="formColor"
                label="Description"
                name="description"
                rows=5
                auto-grow
                required
                ></v-textarea>
            </v-flex>
          </v-layout>

          <v-layout row>
            <v-spacer></v-spacer>
            <v-btn type="submit" :color="formColor" flat>
              {{ actionText }}
            </v-btn>
            <v-btn :color="formColor" flat @click="closeTaskForm();">
              {{ cancelText }}
            </v-btn>
          </v-layout>
        </v-form>
      </v-card-text>
    </v-container>
  </v-card>
</v-dialog>
</template>

<script>
import axios from 'axios';
import { mapActions, mapGetters, mapMutations } from 'vuex';

const NEW_ACTION = 'new';
const EDIT_ACTION = 'edit';

export default {
  name: 'TaskForm',
  mounted() {
    this.$nextTick(() => this.$refs.name.focus());
  },
  computed: {
    ...mapGetters('status', [
      'status',
      'statusType',
    ]),
    ...mapGetters('taskForm', [
      'taskOpenForm',
    ]),
    actionText() {
      return (this.action === NEW_ACTION) ? 'Create' : 'Save';
    },
    actionLink() {
      return (this.action === NEW_ACTION) ? '/insert_task' : '/update_task';
    },
    cancelText() {
      return (this.action === NEW_ACTION) ? 'Close' : 'Cancel';
    },
    useDarkTheme() {
      return this.theme === 'dark';
    },
    formColor() {
      if (this.theme === 'light' && this.formcolor === undefined) {
        return 'info';
      }
      return this.formcolor;
    },
    headerColor() {
      if (this.headercolor === undefined) {
        return (this.theme === 'light') ? 'black' : 'white';
      }
      return this.headercolor;
    },
    headerText() {
      return (this.action === NEW_ACTION) ? 'Create Task' : 'Edit Task';
    },
    openTaskForm: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit('input', val);
      },
    },
  },
  data() {
    return {
      dateMenu: false,
      timeMenu: false,

      nameRules: [
        v => !!v || 'Name is required!',
      ],
      dateRules: [
        () => !!this.info.date || 'Date is required!',
      ],
      timeRules: [
        () => !!this.info.time || 'Time is required!',
      ],
    };
  },
  methods: {
    ...mapMutations('status', [
      'setError',
    ]),
    ...mapMutations('task', [
      'insertTask',
      'deleteTask',
      'updateTask',
      'updateTaskIndex',
    ]),
    ...mapActions('taskForm', [
      'closeTaskForm',
    ]),
    activateDatepicker() {
      setTimeout(() => {
        this.dateMenu = true;
      }, 0);
    },
    activateTimepicker() {
      setTimeout(() => {
        this.timeMenu = true;
      }, 0);
    },
    async saveForm() {
      if (!this.$refs.form.validate()) {
        return;
      }

      const task = {
        index: this.info.index,
        name: this.info.name,
        date: this.info.date,
        time: this.info.time,
        description: this.info.description,
        position: this.info.position,
      };
      if (this.action === NEW_ACTION) {
        this.insertTask(task);
      }

      const resp = await fetch('/', response => response);
      axios.post(this.actionLink, task, {
        headers: { 'X-CSRF-TOKEN': resp.headers.get('X-CSRF-TOKEN') },
      }).then((response) => {
        if (response.data.statusType === 'success') {
          if (this.action === NEW_ACTION) {
            this.updateTaskIndex({
              position: task.position,
              index: parseInt(response.data.index, 10),
            });
          } else {
            this.updateTask(task);
          }
          this.closeTaskForm();
        } else {
          this.deleteTask(task);
          this.setError(response.data.status);
        }
      }).catch((error) => {
        this.setError(error);
      });
    },
  },
  props: {
    info: undefined,
    action: {
      default: undefined,
      type: String,
      validator: value => [NEW_ACTION, EDIT_ACTION].indexOf(value) !== -1,
    },

    // value is used to toggle the dialog
    value: {
      default: undefined,
    },
    theme: {
      default: 'light',
      type: String,
      validator: value => ['light', 'dark'].indexOf(value) !== -1,
    },
    formcolor: {
      default: undefined,
      type: String,
    },
    headercolor: {
      default: undefined,
      type: String,
    },
  },
  watch: {
    openTaskForm() {
      this.$refs.form.resetValidation();
      this.$nextTick(() => this.$refs.name.focus());
    },
  },
};
</script>
