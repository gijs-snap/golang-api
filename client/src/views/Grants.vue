<template>
  <div class="w-full overflow-x-hidden min-h-screen bgcolor">
    <div class="w-screen py-8 sm:py-24 bgpic">
      <div class="w-10/12 mx-auto grid grid-cols-1 sm:grid-cols-2 items-center align-center">
        <div class="text-center sm:text-left">
          <img class="w-3/4 sm:w-1/2 pb-8 mx-auto sm:mx-0" src="../assets/logo.png" />
          <button
            class="w-1/2 mb-8 sm:mb-0 sm:py-4 px-4 bg-transparent hover:bg-white text-white hover:text-black font-semibold border border-white rounded focus:outline-none"
          >
            Back to webiste
          </button>
        </div>
        <h1 class="text-4xl lg:text-5xl text-white text-center sm:justify-self-end">
          Greenhill Seltzers Grants
        </h1>
      </div>
    </div>
    <div class="w-10/12 mx-auto py-12 text-white">
      <div class="text-white">
        <h1 class="font-bold text-4xl pb-4">The Problem</h1>
        <p class="text-lg">
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
          minim veniam, quis nostrud exercitation ullamco laboris nisi ut
          aliquip ex ea commodo consequat. Duis aute irure dolor in
          reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
          pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
          culpa qui officia deserunt mollit anim id est laborum."
        </p>
      </div>

      <div class="text-white text-right py-10">
        <h1 class="font-bold text-4xl pb-4">What is a Grant?</h1>
        <p class="text-lg">
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
          minim veniam, quis nostrud exercitation ullamco laboris nisi ut
          aliquip ex ea commodo consequat. Duis aute irure dolor in
          reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
          pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
          culpa qui officia deserunt mollit anim id est laborum."
        </p>
      </div>

      <div class="text-white py-10 grid grid-cols-1 md:grid-cols-2">
        <div>
          <h1 class="font-bold text-4xl pb-4">Who we are</h1>
          <p class="text-lg pb-4 sm:pb-0">
            "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim
            ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
            aliquip ex ea commodo consequat. Duis aute irure dolor in
            reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
            pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
            culpa qui officia deserunt mollit anim id est laborum."
          </p>
        </div>
        <img
          src="../assets/bg.jpg"
          class="rounded w-3/5 sm:1/2 mx-auto my-auto"
        />
      </div>

      <div class="text-white text-right py-10">
        <h1 class="font-bold text-4xl pb-4">How to Apply?</h1>
        <p class="text-lg">
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
          minim veniam, quis nostrud exercitation ullamco laboris nisi ut
          aliquip ex ea commodo consequat. Duis aute irure dolor in
          reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
          pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
          culpa qui officia deserunt mollit anim id est laborum."
        </p>
      </div>

      <div class="text-white">
        <h1 class="font-bold text-4xl py-10">Inspiration</h1>
        <p class="text-lg">
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
          minim veniam, quis nostrud exercitation ullamco laboris nisi ut
          aliquip ex ea commodo consequat. Duis aute irure dolor in
          reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
          pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
          culpa qui officia deserunt mollit anim id est laborum."
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.bgpic {
  background-image: url("../assets/bg.jpg");
  background-repeat: no-repeat;
  background-position: center;
  background-size: cover;
}

.bgcolor {
  background-color: #1a1a1a;
}
</style>