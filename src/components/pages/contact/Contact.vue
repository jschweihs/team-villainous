<template>
  <div class="content half">
    <h1>Contact</h1>
    <form>
      <input type="text" v-model="name" placeholder="Name" required/>
      <input type="text" v-model="email" placeholder="Email" required/>
      <select v-modoel="category" required>
        <option value="" disabled selected hidden>Category</option>
        <option>Team Request</option>
        <option>Tournament Entry</option>
      </select>
      <textarea
        placeholder="Enter message here..."
        v-model="message"
        requried
      >
      </textarea>
      <input type="submit"/>
    </form>
    <p>Thank you for reaching out to Team Villainous. Please allow our team to respond to your inquiry within 24 hours. You can also contact us on any of our social media.</p>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        name:       '',
        email:      '',
        category:   '',
        message:    '',
        submitted:  false
      };
    },
    methods: {
      submitEmail() {
        // Use axios to submit email
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
