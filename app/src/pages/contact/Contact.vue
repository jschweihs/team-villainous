<template>
  <div class="content half">
    <h1>Contact</h1>
    <form method="post" @submit.prevent="sendEmail">
      <input type="text" v-model="email.name" placeholder="Name" required/>
      <input type="text" v-model="email.address" placeholder="Email" required/>
      <select v-modoel="email.category" required>
        <option value="" disabled selected hidden>Category</option>
        <option>Team Request</option>
        <option>Tournament Entry</option>
      </select>
      <textarea
        placeholder="Enter message here..."
        v-model="email.message"
        requried
      >
      </textarea>
      <input type="submit"/>
    </form>
    <p>Thank you for reaching out to Team Villainous. Please allow our team to respond to your inquiry within 24 hours. You can also contact us on any of our social media.</p>
  </div>
</template>

<script>

  import axios from 'axios';
  
  export default {
    data() {
      return {
        email: {
          name:       '',
          address:    '',
          category:   '',
          message:    ''
        }
      };
    },
    methods: {
      sendEmail() {
        axios.post('http://teamvillainous.com/api/v1/file/send-email', {...this.email})
        .then(res => {
          console.log('submitted successfully');
        })
        .catch(e => console.log(e));
      }
    }
  }
</script>
<style scoped>

  form {
    width: 100%;
    margin: 0 auto;
  }
  
  input, select, textarea {
    width: 100%;
    border-radius: 8px;
    font-size: 24px;
    padding: 10px;
    margin: 10px 0;
    font-family: 'Nixie One', cursive;
    border: 0;
    box-sizing: border-box;
  }

  textarea {
    height: 200px;
  }

  input[type=submit] {
    background-color: #ffc200;
    cursor: pointer;
    color: white;
    font-weight: bold;
    height: 60px;
  }

  select:required:invalid {
    color: grey;
  }

  select option:not(:disabled){
  color: black;
}

</style>
