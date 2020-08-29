<template>
  <div class="flex justify-center mb-5">
    <user-avatar v-if="profile" :url="profile.avatar_url" />
    <div class="pl-6 flex flex-col justify-center">
      <h1>{{profile.name}}</h1>
      <h2><a :href="profile.html_url">Github</a></h2>
      <h2 v-if="profile.twitter_username">
        <a :href="'https://twitter.com/' + profile.twitter_username">
          Twitter
        </a>
      </h2>
      <h2 v-if="profile.company">Works: {{profile.company}}</h2>
      <h2 v-if="profile.company">Location: {{profile.location}}</h2>
    </div>
  </div>
  <config-info v-if="vimconfig" :vimconfig="vimconfig" />
</template>

<script>
import ConfigInfo from '../components/config-info.vue';
import UserAvatar from '../components/user-avatar.vue';

async function fetchProfile(user) {
  try {
    const url = 'http://localhost:3001/api/github/user?login=' + user;
    const res = await window.fetch(url);
    return await res.json();
  } catch (err) {
    console.log(err);
  }
}

async function fetchVimrc({ user, repo }) {
  try {
    const url = `http://localhost:3001/api/github/vimrc`;
    const res = await window.fetch(url + `?login=${user}&repo=${repo}`);
    return await res.json();
  } catch (err) {
    console.log(err);
  }
}

export default {

  props: {
    github: {
      type: Object,
      required: true
    }
  },

  components: {
    ConfigInfo,
    UserAvatar
  },

  async setup(props) {
    const profile = await fetchProfile(props.github.user);
    const vimconfig = await fetchVimrc(props.github);
    return { vimconfig, profile };
  }

};
</script>
